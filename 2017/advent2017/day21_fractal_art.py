"""
http://adventofcode.com/2017/day/21

I couldn't find a matrix library with good indexing and
concatenation functionality. This is a good opportunity
to become familiar with numpy.
"""

import numpy as np

def str_to_mat(pattern):
	return np.array([[0 if e == "." else 1 for e in row]
					for row in pattern.split("/")])

def rotations_and_flips(pattern):
	yield pattern # original
	yield np.rot90(pattern) # rotate 90
	yield np.rot90(pattern, 2) # rotate 180
	yield np.rot90(pattern, 3) # rotate 270
	yield np.flip(pattern, 0) # flip rows
	yield np.rot90(np.flip(pattern, 0), 1) # flip rows and rot 90
	yield np.rot90(np.flip(pattern, 0), 2) # flip rows and rot 180
	yield np.rot90(np.flip(pattern, 0), 3) # flip rows and rot 270

class Rule:
	def __init__(self, rule_s):
		"""
		rule_s is like ../.# => ##./#../...
		The components of rule_s are
			<pattern> => <replacement>
		Both pattern and replacement are converted to
		numpy arrays.
		"""
		pattern, replacement = [str_to_mat(x) for x in rule_s.split(" => ")]
		self.pattern = pattern
		self.replacement = replacement
		self.size = self.pattern.shape[0]

	def match(self, pattern):
		for option in rotations_and_flips(pattern):
			if np.array_equal(self.pattern, option):
				return np.copy(self.replacement)

		return None

class Ruleset:
	def __init__(self, rules_s):
		"""
		rules_s is a string with rule strings on separate lines.
		"""
		self.rules = [Rule(line) for line in rules_s.split("\n")]

	def match(self, pattern):
		for rule in self.rules:
			replacement = rule.match(pattern)
			if replacement is not None:
				return replacement

		return None

if __name__ == "__main__":
	pattern = str_to_mat(".#./..#/###")

	with open("day21_input.txt", "r") as f:
		rules = Ruleset(f.read())

	for i in range(18):
		print(i, np.sum(pattern), pattern.shape)
		nrow, ncol = pattern.shape
		if nrow % 2 == 0:
			pattern = np.vstack([np.hstack([rules.match(col) for col in np.split(row, ncol / 2, 1)]) for row in np.split(pattern, nrow / 2, 0)])
		else:
			pattern = np.vstack([np.hstack([rules.match(col) for col in np.split(row, ncol / 3, 1)]) for row in np.split(pattern, nrow / 3, 0)])

	print(np.sum(pattern))
