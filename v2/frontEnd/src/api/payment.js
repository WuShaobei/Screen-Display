///
 // @Author: WuShaobei
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-13 15:21:20
 // @FilePath: /Screen-Display/v2/frontEnd/src/api/payment.js
 // @Description: 
 ///

 import axios from 'axios'
 
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