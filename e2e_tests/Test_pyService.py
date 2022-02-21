import unittest

from BaseTest import BaseTest


class Test_pyService(BaseTest):
    PY_SERVICE_URL = "http://localhost:8000/health"

    def test_pyService(self):
        self.isBackendAlive(self.PY_SERVICE_URL,"python service")


if __name__ == '__main__':
    unittest.main()
