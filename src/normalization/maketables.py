#!/usr/bin/env python


import string
import httplib
import sys
import re

FCodePoint = 0
FName = 1
FGeneralCategory = 2
FCanonicalCombiningClass = 3
FBidiClass = 4
FDecompositionType = 5
FDecompositionMapping = 6
FNumericType = 7
FNumericValue = 8
FBidiMirrored = 9
FUnicode1Name = 10
FISOComment = 11
FSimpleUppercaseMapping = 12
FSimpleLowercaseMapping = 13
FSimpleTitlecaseMapping = 14
NumField = 15

MaxChar = 0x10FFFF # anything above this shouldn't exist


exclusions = []
canonical = dict() # sortedList
compatibility = dict() #sortedList
combiningClasses = dict() #sortedList

compatibilityKeys = dict()
compatibilityMappings = []

firstMap = dict()
secondMap = dict()



composeLookupMax = 0
singleFirstComposition = dict()
singleSecondComposition = dict()
complexComposition = dict()

def smartSort(item):
	print item
	if type(item).__name__=='str':
		return int(item, 16)
	return int(item)

def max(a, b):
	if a>b:
		return a
	return b

def stripComment(line):
	return string.split(line, "#")[0]

def decompose(input, mappings):
	out = ""
	c = input.split(" ")
	
	for d in c:
		if d in mappings:
			if len(out) > 0:
				out += " "
			out += decompose(mappings[d], mappings)
		else:
			if len(out) > 0:
				out += " "
			out += d
			
	return out

def isCompatibilityMapping(input):
	return len(input) > 0 and input[0] == "<"
	
def stripCompatibilityTag(input):
	out = input.split("> ")[1]
	out = out.strip()
	return out
	
def loadExclusions():
	conn = httplib.HTTPConnection("www.unicode.org")
	conn.request("GET", "/Public/5.2.0/ucd/CompositionExclusions.txt")
	resp = conn.getresponse()
	
	if resp.status != 200:
		sys.exit("Error downloading CompositionExclusions.txt: Status "+resp.status+" "+resp.reason)

	input = resp.read().splitlines()
	
	for line in input:
		line = stripComment(line)
		line = line.strip()
		# Hack: The codepoints are represented as either 4 or 5 letter hex value. 
		# This zerofills the 4 letter ones to make things easier to sort.
		line = re.sub(r'((?<![A-F0-9])([A-F0-9]{4})(?![A-F0-9]))', r'0\1', line)
		
		if line != "":
			exclusions.append(line)
		
	resp.close()

def loadUnicodeData():
	conn = httplib.HTTPConnection("www.unicode.org")
	conn.request("GET", "/Public/5.2.0/ucd/UnicodeData.txt")
	resp = conn.getresponse()
	
	if resp.status != 200:
		sys.exit("Error downloading UnicodeData.txt: Status "+resp.status+" "+resp.reason)
		
	
	input = resp.read().splitlines()
	
	for line in input:
		line = stripComment(line)
		# Hack: The codepoints are represented as either 4 or 5 letter hex value. 
		# This zerofills the 4 letter ones to make things easier to sort.
		line = re.sub(r'((?<![A-F0-9])([A-F0-9]{4})(?![A-F0-9]))', r'0\1', line)
		line = line.strip()
		
		if line != "":
			
			f = line.split(";")
			#if len(f[FCodePoint]) == 4:
			#line = re.sub(r'((?<![A-F0-9])([A-F0-9]{4})(?![A-F0-9]))', r'0\1', line)
			#f = line.split(";")
			
			if len(f[FDecompositionType]) == 5:
				exclusions.append(f[FCodePoint])
				
			
			if f[FDecompositionType] != "":
				if isCompatibilityMapping(f[FDecompositionType]):
					compatibility[f[FCodePoint]] = stripCompatibilityTag(f[FDecompositionType])
				else:
					compatibility[f[FCodePoint]] = f[FDecompositionType]
					if f[FCodePoint] not in exclusions:
						canonical[f[FCodePoint]] = f[FDecompositionType]
			if f[FCanonicalCombiningClass] != "0":
				combiningClasses[int(f[FCodePoint], 16)] = f[FCanonicalCombiningClass]
		
	resp.close()
	

def applyCompatibilityMappings():
	while True:
		replaced = False
		
		for key, value in sorted(compatibility.items()):
			d = decompose(value, compatibility)
			if d != value:
				replaced = True
				compatibility[key] = d

		if not replaced:
			break


			
def eliminateDuplicateMappings():
		for key, value in sorted(compatibility.items()):
			index = -1 
			try:
				index = compatibilityMappings.index(value)
			except ValueError:
				pass
				
			if index == -1:
				index = len(compatibilityMappings)
				compatibilityMappings.append(value)
				
			compatibilityKeys[key] = index
			


def createCompositionTables():
	##print canonical
	#sys.exit()
	
	global composeLookupMax
	
	for key, value in sorted(canonical.items()):
		s = value.split(" ")
		
		if len(s) == 2:
			# If both characters have the same combining class, they
			# won't be combined (in the sequence AB, B is blocked from
			# A if both have the same combining class)   
			
			cc1 = None 
			try:
				cc1 = combiningClasses[int(s[0], 16)]
			except KeyError:
				pass
				
			cc2 = None
			try:
				combiningClasses[int(s[1], 16)]
			except KeyError:
				pass
				
				
			if cc1 != None or (cc1 != None and cc1==cc2):
				# ignore this composition
				del canonical[key]
				continue
			
			if s[0] in firstMap:
				c = int(firstMap[s[0]])
				firstMap[s[0]] = c + 1
			else:
				firstMap[s[0]] = 1
				
			if s[1] in secondMap:
				c = int(secondMap[s[1]])
				secondMap[s[1]] = c+1
			else:
				secondMap[s[1]] = 1
		elif len(s) > 2:
			sys.exit("Wrong canonical mapping for "+ key)
				
	
	for key, value in sorted(canonical.items()):
		s = value.split(" ")
		
		if len(s) == 2:
			
			first = 0
			if s[0] in firstMap:
				first = int(firstMap[s[0]])
				
			second = 0
			if s[1] in secondMap:
				second = int(secondMap[s[1]])
				
			if first == 1:
				#global composeLookupMax
				singleFirstComposition[s[0]] = [s[1], key]
				composeLookupMax = max(composeLookupMax, int(s[0],16))
			elif second == 1:
				#global composeLookupMax
				singleSecondComposition[s[1]] = [s[0], key]
				composeLookupMax = max(composeLookupMax, int(s[1],16))
			else:
				if s[0] in complexComposition:
					m = complexComposition[s[0]]
					if s[1] in m:
						sys.exit("Ambiguous canonical mapping for " + s[0])
					m[s[1]] = key
				else:
					m = dict()
					m[s[1]] = key
					complexComposition[s[0]] = m
					
				#global composeLookupMax
				composeLookupMax = max(composeLookupMax, int(s[0],16))
				composeLookupMax = max(composeLookupMax, int(s[1],16))

def printAll():
	print "// Copyright 2010 Hannes Baldursson. All rights reserved."
	print "// Use of this source code is governed by a BSD-style"
	print "// license that can be found in the LICENSE file.\n"
	
	print "// Do Not Edit !!!!"
	print "// This file is generated automatically!\n"
	
	print "// This file is part of go-idn \n"
	
	print "package normalization\n"
	
	# dump combining classes
	
	print "var _CombiningClass_c = [][]int {"

	index = ""
	count = 0
	
	# Where do these numbers in 'range' come from?
	for i in range(0,700):
		empty = True
		
		page = ""
		page += "	[]int { /* Page " + str(i) + " */"
		
		for j in range(0,256):
			c = ((i<<8)+j)
			cc = None
			
			try:
				cc = combiningClasses[c]
			except:
				pass
			
			if 0 == (j & 31):
				page += "\r\n		"
			if cc == None:
				page += "0, "
			else:
				page += cc + ", "
				empty = False
		
		page += "\r\n	},"
		index += "	"
		
		if not empty:
			print page
			index += str(count)
			index += ",\r\n"
			count = count+1
		else:
			index += "-1,\r\n"
	print "	}\r\n"
	
	print "var _CombiningClass_i = []int {"
	print index
	#print "	};"
	print "}\r\n"
	
	print "var _Decomposition = map[int][]int {"
	for key, index in sorted(compatibilityKeys.items()):
		line = "\t"
		line += "0x"+str(key) + " : []int {"
		maps = compatibilityMappings[index].split(" ")
		for i in range(0,len(maps)):
			if i == len(maps)-1:
				line += "0x"+maps[i]+"},"
			else:
				line += "0x"+maps[i]+", "
		print line
	print "}"

		
		
	# dump canonical composition
	index = 0
	indices = dict()
	for key in sorted(complexComposition.keys()):
		indices[int(key,16)] = index
		index = index + 1
	#	print index
	
	
	
	multiSecondStart = index
	print "	/* jagged array */"
	print "var _Composition_multiFirst = [][]int {"
		
	for s0,m in sorted(complexComposition.items()):
		
		line = dict()
		strline =""
		maxIndex = 1
		
		
		
		
		for s1, k in sorted(m.items()):
			s1i = int(s1, 16)
			
			if s1i not in indices:
				indices[s1i] = index
				index = index + 1
			line[indices[s1i]] = k
			maxIndex = max(maxIndex, indices[s1i])
		
		strline += "	[]int {"
		for j in range(multiSecondStart, maxIndex+1): # ATHUGA!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
			
			if j == maxIndex:
				if j in line:
					strline += "0x"+line[j]+"}, "
				else:
					strline += "0x00000}, "
			else:
				if j in line:
					strline += "0x"+line[j]+", "
				else:
					strline += "0x00000, "

		print strline
	print "}"
	
	singleFirstStart = index
	
	print "var _Composition_singleFirst = [][]int {"
	for key, value in sorted(singleFirstComposition.items()):
		print "	[]int {0x"+value[0]+", 0x"+value[1]+"},"
		
		indices[int(key,16)] = index
		index = index +1
	print "}"
	
	singleSecondStart = index
	
	print "var _Composition_singleSecond = [][]int {"
	for key, value in sorted(singleSecondComposition.items()):
		print "	[]int {0x"+value[0]+", 0x"+value[1]+"},"
		
		indices[int(key,16)] = index
		index = index+1
	print "}\n"	
	
	print "var _Composition_multiSecondStart int = "+ str(multiSecondStart)
	print "var _Composition_singleFirstStart int = "+ str(singleFirstStart)
	print "var _Composition_singleSecondStart int = "+ str(singleSecondStart)
		
	print "var _Composition_composePage = []int {"
	pageCount = 0
	compositionPages = ""
	j=0
	
	while  (j*256) < (composeLookupMax+255):
		empty = True
		page = ""
		
		for k in range(0, 256):
			if k% 16 == 0:
				page += "\r\n	"
			if (j *256+k) in indices:
				page += str(indices[j*256+k])
				page += ', '
				empty = False
			else:
				page += "-1, "
			
		if empty:
			print "	-1,"
		else:
			print "	" + str(pageCount) + ","
			compositionPages += "	[]int {"
			compositionPages += page
			compositionPages += "\r\n	},\r\n"
			pageCount = pageCount +1
			
		j = j+1
	print "}"
	
	print "var _Composition_composeData = [][]int {"
	print compositionPages
	print "}"
	
			
	
	
	
	
	
if __name__ == "__main__":
	loadExclusions()
	loadUnicodeData()
	applyCompatibilityMappings()
	eliminateDuplicateMappings()
	createCompositionTables()
	printAll()
	
	'''print compatibility
	print compatibilityKeys
	print compatibilityMappings'''

	
	