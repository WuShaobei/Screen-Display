import axios from 'axios'

function postApiSelectChineseCateringPayment(callback){
    
    axios.post(
        // TODO 范围查询
        "http://127.0.0.1:1432/api/v1/select/ChineseCateringPayment?BeginYear=2018&EndYear=2021"
    ).then(function(res){
        if (res.data.Code == 0) {
            callback({Data:res.data.Data})
        }
    })
}

export{
    postApiSelectChineseCateringPayment
}