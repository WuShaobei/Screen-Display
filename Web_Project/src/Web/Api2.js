import axios from 'axios'
function getApi2Data (callback){
    axios.post(
        "http://127.0.0.1:1432/api/v1/api2?BeginYear=2010&EndYear=2021"
    ).then( function(res){
        if (res.data.Code == 0){
            callback({
                Data: {
                    title:{
                        text:"2010年-2021年中国餐饮线上订单走势",
                    },
                    xAxis: {
                        type: 'category',
                        data: res.data.AllTime
                    },
                    yAxis: {
                        type: 'value'
                    },
                    series: [
                        {
                            data: res.data.AllOrderAmount,
                            type: 'line',
                            smooth: true,
                            color:'#5555AA',
                        }
                    ]
                }
            })
        }
    }).catch(function(err){
        console.log(err)
    })
}
export{
    getApi2Data,
}