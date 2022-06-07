# Example
import unittest
from addsub import add, subtract

class TestBasicMethods(unittest.TestCase):

    def test_add(self):
        self.assertEqual(add(2,3), 6)

    def test_subtract(self):
        self.assertEqual(subtract(5,3), 3)

if __name__ == '__main__':
    unittest.main()