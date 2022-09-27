import axios from 'axios'

function getApi3Data(callback){
    
    axios.post(
        // TODO 范围查询
        "http://127.0.0.1:1432/api/v1/api3?BeginYear=2018&EndYear=2021"
    ).then(function(res){
        if (res.data.Code == 0) {
            callback({Data:res.data.Data})
        }
    })
}

export{
    getApi3Data
}