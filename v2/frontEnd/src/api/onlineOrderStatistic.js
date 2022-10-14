///
 // @Author: WuShaobei
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-13 15:20:58
 // @FilePath: /Screen-Display/v2/frontEnd/src/api/onlineOrderStatistic.js
 // @Description: 
 ///
 import axios from 'axios'
 
 function postAmountByYearAndMonthFromOnlineOrderStatisticApi(beginYear, endYear, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/data/postAmountByYearAndMonthFromOnlineOrderStatistic", {
            beginYear : beginYear,
            endYear : endYear
        }
    ).then( function(res){
        if (res.data.code == 0){
            let data = {
                title:{
                    text:"2010年-2021年中国餐饮线上订单走势",
                },
                xAxis: {
                    type: 'category',
                    data: res.data.data.allTime
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: res.data.data.allAmount,
                        type: 'line',
                        smooth: true,
                        color:'#5555AA',
                    }
                ]
            }
            callback({
                code : res.data.code,
                data : data
            })
        } else {
            callback({code : res.data.code})
        }
    }).catch(function(err){
        console.log(err)
    })
 }

 export{
    postAmountByYearAndMonthFromOnlineOrderStatisticApi
 }