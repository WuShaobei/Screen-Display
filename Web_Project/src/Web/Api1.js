import axios from "axios";
function getApi1Data(callback){    
    axios.post(
        "http://127.0.0.1:1432/api/v1/api1?BeginYear=2014&EndYear=2021"
    ).then(function(res){
        if (res.data.Code == 0){
            let colors = ['#559955', '#FF9900'];
            let xData = [];
            let yData1 = [];
            let yData2 = [];
            for(let i = 0; i <= 2021 - 2014; i ++) {
                xData.push(res.data.Data[i].Year)
                yData1.push(res.data.Data[i].TotalAmount)
                yData2.push(res.data.Data[i].TotalAmountPercentage)
            }
            
            let data = {
                title : {
                    text:"2014年-2021年中国餐饮市场规模"
                },
                color: colors,
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                    type: 'cross'
                    }
                },
                grid: {
                    right: '25%'
                },
                xAxis: [
                    {
                    type: 'category',
                    axisTick: {
                        alignWithLabel: true
                    },
                    // prettier-ignore
                    // data: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
                        data: xData
                    }
                ],
                yAxis: [
                    {
                    type: 'value',
                    name: '市场份额(万亿)',
                    position: 'right',
                    alignTicks: true,
                    offset: 80,
                    axisLine: {
                        show: true,
                        lineStyle: {
                        color: colors[0]
                        }
                    },
                    axisLabel: {
                        formatter: '{value} 万亿'
                    }
                    },
                    {
                    type: 'value',
                    name: '百分比',
                    position: 'left',
                    alignTicks: true,
                    axisLine: {
                        show: true,
                        lineStyle: {
                        color: colors[1]
                        }
                    },
                    axisLabel: {
                        formatter: '{value} %'
                    }
                    }
                ],
                series: [
                    {
                    name: '市场份额(万亿)',
                    type: 'bar',
                    yAxisIndex: 0,
                    // data: [
                    //     2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3
                    // ]
                    data : yData1
                    },
                    {
                    name: '百分比',
                    type: 'line',
                    yAxisIndex: 1,
                    // data: [2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 6.2]
                    data: yData2
                    }
                ]
            
            
            }

            callback( {Data : data} )
        }
    }).catch(function(err){
        console.log(err)
    })
}

export{
    getApi1Data
}