import unittest

from BaseTest import BaseTest


class Test_goService(BaseTest):
    GO_SERVICE_URL = "http://localhost:3000/health"

    def test_goService(self):
        self.isBackendAlive(self.GO_SERVICE_URL, "go service")


if __name__ == '__main__':
    unittest.main()
