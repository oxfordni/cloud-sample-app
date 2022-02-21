import time
import unittest

import requests
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options

class BaseTest(unittest.TestCase):
    FRONTEND_URL = "http://localhost:3001"
    COLOR_GREEN = "rgba(82, 196, 26, 1)"
    COLOR_RED = "rgba(255, 77, 79, 1)"

    def isBackendAlive(self, url,xpath):
        print("\nChecking service alive: "+url)
        # Check service status
        status = False
        try:
            response = requests.get(url)
        except requests.exceptions.RequestException as e:
            status = False
        else:
            print(response.status_code)
            if response.status_code == 200:
                message = response.json()
                print(message)
                for key, value in message.items():
                    if key == 'alive' and value == 'ok':
                        print(value)
                        status = True
                    else:
                        status = False
            else:
                status = False

        # Get color in UI
        chrome_options = Options()
        chrome_options.add_argument('--headless')
        chrome_options.add_argument('--no-sandbox')
        chrome_options.add_argument('--disable-dev-shm-usage')
        driver = webdriver.Chrome(options=chrome_options)
        driver.implicitly_wait(20)
        driver.get(self.FRONTEND_URL)
        driver.maximize_window()
        time.sleep(20)
        element_pyservice = driver.find_element(By.XPATH, "//span[contains(text(), '"+xpath+"')]")
        color_pyservice = element_pyservice.value_of_css_property("color")

        if status:
            # Verify UI is GREEN
            self.assertEqual(self.COLOR_GREEN, color_pyservice, "UI is not in GREEN color while service is UP")
            print(color_pyservice)
        else:
            # Verify UI is RED
            self.assertEqual(self.COLOR_RED, color_pyservice, "UI is not in RED color while service is DOWN")

        driver.close()


if __name__ == '__main__':
    unittest.main()
