from selenium import webdriver
import time
from flask import Flask, jsonify, request
from concurrent.futures import ThreadPoolExecutor

app = Flask(__name__)


# DownEogData 下载eog data
class DownEogData:
    def __init__(self, m):
        # self.data_url = "https://eogdata.mines.edu/pages/download_dnb_composites_iframe.html"
        options = webdriver.ChromeOptions()
        options.add_argument('--no-sandbox')
        options.add_argument('--disable-dev-shm-usage')
        options.add_argument('--headless')
        options.add_argument('blink-settings=imagesEnabled=false')
        options.add_argument('--disable-gpu')
        prefs = {'download.default_directory': m['download_dir']}
        options.add_experimental_option('prefs', prefs)
        self.driver = webdriver.Chrome(chrome_options=options)
        if len(m["data"]) > 1:
            self.first_href = m["data"][0]["href"]
            self.data = m["data"][1:]
        elif len(m["data"]) == 1:
            self.first_href = m["data"][0]["href"]
            self.data = []

    # 登录下载数据
    def download(self):
        self.driver.get(self.first_href)
        username = self.driver.find_element_by_xpath('//*[@id="username"]')
        username.send_keys("2685366884@qq.com")
        password = self.driver.find_element_by_xpath('//*[@id="password"]')
        password.send_keys("qwerdf123")
        rememberMe = self.driver.find_element_by_xpath('//*[@id="rememberMe"]')
        rememberMe.click()
        login = self.driver.find_element_by_xpath('//*[@id="kc-login"]')
        login.click()
        print("下载文件: "+self.first_href)
        for item in self.data:
            time.sleep(3)
            self.driver.get(item["href"])
            print("下载文件: "+item["href"])


def start_download(req):
    e = DownEogData(req)
    e.download()


@app.route("/download", methods=['POST'])
def download():
    executor = ThreadPoolExecutor(1)
    req = request.get_json()
    try:
        executor.submit(start_download, req)
        return jsonify({"msg": "success"})
    except Exception as e:
        return jsonify({"msg": e})


if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True, port='8033')
