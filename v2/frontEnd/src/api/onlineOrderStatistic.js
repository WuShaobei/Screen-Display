///
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-14 17:50:47
 // @FilePath: /frontEnd/src/api/onlineOrderStatistic.js
 // @Description: 网上订单类接口
 ///
 import axios from 'axios'
 
 /******* 
  * @description: post "http://127.0.0.1:1432/api/data/postAmountByYearAndMonthFromOnlineOrderStatistic" 并返回折线图数据
  * @param {*} beginYear
  * @param {*} endYear
  * @param {*} callback
  * @return {*}
  */
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