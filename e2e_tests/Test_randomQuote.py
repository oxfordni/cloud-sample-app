import unittest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
import time

class Test_randomQuote(unittest.TestCase):
    ALERT_MSG = "Random Quote"

    def test_quotescheck(self):
        s = Service(r"C:\Users\User\Downloads\chromedriver.exe")
        self.driver = webdriver.Chrome(service=s)
        self.driver.implicitly_wait(20)
        self.driver.get("http://localhost:3001")
        self.driver.maximize_window()
        time.sleep(10)
        # Check random quote is displayed while initial load
        alertMessage = self.driver.find_element(By.CLASS_NAME, "ant-alert-message").text
        self.assertEqual(alertMessage, self.ALERT_MSG, "Random quote not displayed")
        alertDes = self.driver.find_element(By.CLASS_NAME, "ant-alert-description").text
        time.sleep(15)
        # Check random quote is displayed after refresh
        self.driver.refresh()
        alertMessageRefresh = self.driver.find_element(By.CLASS_NAME, "ant-alert-message").text
        self.assertEqual(alertMessage, alertMessageRefresh, "Random quote not displayed after refresh")
        alertDesRefresh = self.driver.find_element(By.CLASS_NAME, "ant-alert-description").text
        print("Initial alert message: ",alertMessage)
        print("Initial quote: ", alertDes)
        print("Alert message after refresh: ",alertMessageRefresh)
        print("Quote after refresh: ", alertDesRefresh)
        # Check initial and refreshed quotes
        self.assertNotEqual(alertDes, alertDesRefresh, "UI showed same quote after refresh")

if __name__ == "__main__":
    unittest.main()
