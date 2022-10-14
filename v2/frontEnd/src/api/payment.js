///
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-14 17:50:54
 // @FilePath: /frontEnd/src/api/payment.js
 // @Description: 薪资类数据
 ///

 import axios from 'axios'
 
 /******* 
  * @description: post "http://127.0.0.1:1432/api/data/postAvgSalaryAndCountsByYearFromPayment" 并返回数据
  * @param {*} beginYear
  * @param {*} endYear
  * @param {*} callback
  * @return {*}
  */ 
 function postAvgSalaryAndCountsByYearFromPaymentApi(beginYear, endYear, callback) {
    axios.post(
        "http://127.0.0.1:1432/api/data/postAvgSalaryAndCountsByYearFromPayment", {
            beginYear : beginYear,
            endYear : endYear
        }
    ).then(function(res){
        if (res.data.code == 0 ) {            
            callback({code: 0, data:res.data.data})
        } else {
            callback({code: res.data.code})
        }
    }).catch(function(err) {
        console.log(err)
    })
 }

 export{
    postAvgSalaryAndCountsByYearFromPaymentApi
 }