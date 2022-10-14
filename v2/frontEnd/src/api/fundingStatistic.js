///
 // @Author: WuShaobei
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-13 15:21:13
 // @FilePath: /Screen-Display/v2/frontEnd/src/api/fundingStatistic.js
 // @Description: 
 ///
// 

import axios from 'axios'

function getAllDataByFromFundingStatisticApi(callback) {
    axios.get(
        "http://127.0.0.1:1432/api/data/getAllDataByFromFundingStatistic"
    ).then(function(res){
        if ( res.data.code == 0 ) {
            callback({code : res.data.code, data : res.data.data})
        } else {
            callback({code : res.data.code})
        }
    }).catch(function(err){
        console.log(err)
    })
}

export {
    getAllDataByFromFundingStatisticApi   
}