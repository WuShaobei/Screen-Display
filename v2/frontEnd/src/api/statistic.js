///
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-14 17:54:15
 // @FilePath: /frontEnd/src/api/statistic.js
 // @Description: 市场规模相关接口
 ///

 import axios from 'axios'
 

 /******* 
  * @description: post "http://127.0.0.1:1432/api/data/postAmountAndPercentageByYearFromStatistic" 并返回柱状图数据
  * @param {*} beginYear
  * @param {*} endYear
  * @param {*} callback
  * @return {*}
  */
 function postAmountAndPercentageByYearFromStatisticApi(beginYear, endYear, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/data/postAmountAndPercentageByYearFromStatistic", {
            beginYear : beginYear,
            endYear : endYear
        }

    ).then(function(res){
        if (res.data.code == 0){
            let colors = ['#559955', '#FF9900'];
            let xData = [];
            let yData1 = [];
            let yData2 = [];
            for(let i = 0; i <= endYear - beginYear; i ++) {
                xData.push(res.data.data[i].year)
                yData1.push(res.data.data[i].amount)
                yData2.push(res.data.data[i].percentage)
            }
            
            let data = {
                title : {
                    text:beginYear + "年-" + endYear + "年中国餐饮市场规模（万亿）"
                },
                color: colors,
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'cross'
                    }
                },
                grid: {
                    right: '15%',
                    left: '15%'
                },
                xAxis: [
                    {
                    type: 'category',
                    offset: 10,
                    axisTick: {
                        alignWithLabel: true
                    },
                        data: xData
                    }
                ],
                yAxis: [
                    {
                    type: 'value',
                    name: '市场份额',
                    position: 'right',
                    alignTicks: true,
                    offset: 10,
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
                    offset: 10,
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
                    name: '市场份额',
                    type: 'bar',
                    yAxisIndex: 0,
                    data : yData1
                    },
                    {
                    name: '百分比',
                    type: 'line',
                    yAxisIndex: 1,
                    data: yData2
                    }
                ]
            }

            callback( {data : data} )
        }
    }).catch(function(err){
        console.log(err)
    })
 }

 export{
    postAmountAndPercentageByYearFromStatisticApi 
 }