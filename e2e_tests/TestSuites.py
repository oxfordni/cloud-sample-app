import unittest
from Test_randomQuote import Test_randomQuote
from Test_goService import Test_goService
from Test_pyService import Test_pyService

tc1=unittest.TestLoader().loadTestsFromTestCase(Test_randomQuote)
tc2=unittest.TestLoader().loadTestsFromTestCase(Test_pyService)
tc3=unittest.TestLoader().loadTestsFromTestCase(Test_goService)

EndtoendTestSuite=unittest.TestSuite([tc1,tc2,tc3])
unittest.TextTestRunner().run(EndtoendTestSuite)
