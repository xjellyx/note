from frame.Product_data import ProductData
from frame.chart import  Chart


if __name__=="__main__":
    req = ProductData(accessToken="b0de8bdb-4044-4bb8-a6d8-28321a82fd83")
    cht = Chart(title='st')
    param = {"codes": ["20170011"], "beginTime": "2017-11-11", "endTime": "2020-5-27", }
    result = req.getDataChartBar(params=param)
    x_value = []
    y_value = []
    for v in result['data']['20170011']['data']:
        x_value.append(v['group_by'])
        y_value.append(v['volume'])

    print(x_value)
    print(y_value)
    cht.setSeries({'name':'bar','type':'bar','data':y_value})
    cht.setXAxis({'type': 'category','data':x_value})
    cht.setYAxis()
    data = cht.generateChart()
    print(data)
    print(req.sendMsg(params={"message":data,"accessToken":req.accessToken}))


