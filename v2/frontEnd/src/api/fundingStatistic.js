///
 // @Date: 2022-10-10 17:36:47
 // @LastEditTime: 2022-10-14 17:50:06
 // @FilePath: /frontEnd/src/api/fundingStatistic.js
 // @Description: 融资类接口
 ///
 // 
 
 /******* 
  * @description: get "http://127.0.0.1:1432/api/data/getAllDataByFromFundingStatistic" 并返回
  * @param {*} callback
  * @return {*}
  */ 
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