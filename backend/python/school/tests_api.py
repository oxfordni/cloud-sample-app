from rest_framework import status
from rest_framework.test import APITestCase, APIClient

class SimpleAPITestCase(APITestCase):
    @classmethod
    def setUpClass(cls):
        super().setUpClass()
        cls.client = APIClient()
        cls.index_url = ''
        cls.health_url = '/health'

    def test_index(self):
        """Index returns 204 No Content."""
        response = self.client.get(self.index_url)
        self.assertEqual(response.status_code, status.HTTP_204_NO_CONTENT)

    def test_health(self):
        """Health returns 200 OK."""
        response = self.client.get(self.health_url)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(str(response.data), "{'alive': 'ok'}")
