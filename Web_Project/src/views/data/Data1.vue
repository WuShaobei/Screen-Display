<template>
    <div>
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="dataView">
        
            <div class="content">
            
                <div class="leftView">
                    <div class="data1">
                        <div class="leftBox"><div class="text">开始年份:</div></div>
                        <div class="rightBox"><div class="text">
                            <!-- <input v-model="BeginYear" placeholder="开始年份(包含)" id="input" />  -->
                            {{BeginYear}}
                        </div></div>
                    </div>
                    <div class="data2">
                        <div class="leftBox"><div class="text">结束年份:</div></div>
                        <div class="rightBox"><div class="text">
                            {{EndYear}}
                            <!-- <input v-model="EndYear" placeholder="结束年份(包含)" id="input" />  -->
                        </div></div>
                    </div>      
                    
                    <div class="data3"><div class="text">餐饮就业人数行业平均报酬</div></div>
                    <div v-if="showLeftValue" class="data4"><div class="text">
                        <div class="table1">
                            <div class="tableLeft"><div class="textTable">时间</div></div>
                            <div class="tableCenter"><div class="textTable">平均报酬</div></div>
                            <div class="tableRight"><div class="textTable">就业人数</div></div>
                        </div>
                        <div class="table2">
                            <div class="tableLeft"><div class="textTable">{{leftValue[0].Year}}</div></div>
                            <div class="tableCenter"><div class="textTable">{{leftValue[0].AvgSalary}}</div></div>
                            <div class="tableRight"><div class="textTable">{{leftValue[0].CountSalary}}</div></div>
                        </div>
                        <div class="table3">
                            <div class="tableLeft"><div class="textTable">{{leftValue[1].Year}}</div></div>
                            <div class="tableCenter"><div class="textTable">{{leftValue[1].AvgSalary}}</div></div>
                            <div class="tableRight"><div class="textTable">{{leftValue[1].CountSalary}}</div></div>
                        </div>
                        <div class="table4">
                            <div class="tableLeft"><div class="textTable">{{leftValue[2].Year}}</div></div>
                            <div class="tableCenter"><div class="textTable">{{leftValue[2].AvgSalary}}</div></div>
                            <div class="tableRight"><div class="textTable">{{leftValue[2].CountSalary}}</div></div>
                        </div>
                        <div class="table5">
                            <div class="tableLeft"><div class="textTable">{{leftValue[3].Year}}</div></div>
                            <div class="tableCenter"><div class="textTable">{{leftValue[3].AvgSalary}}</div></div>
                            <div class="tableRight"><div class="textTable">{{leftValue[3].CountSalary}}</div></div>
                        </div>
                    </div></div>
                    
                </div>

                <div class="rightView">
                    <div class="top">
                        <e-charts :option="opApi1"/>
                    </div>
                    <div class="down">
                        <e-charts :option="opApi2"/>
                    </div>
                </div>
                    
            </div>


            <div class="indexes">
                <h1 id="h1">
                    <div class="data0"><div class="textIndex">索引页</div></div>
                    <div class="data1"><div class="textIndex">亲爱的:{{getIdentity()}}</div></div>
                    <div class="data2"><div class="textIndex">用户名:{{Username}}</div></div>
                    <div class="data3"><div class="textIndex">真实姓名:{{RealName}}</div></div>
                    <div class="data4"><div class="textIndex"></div></div>
                </h1>
            </div>
        
        </div>
    </div>
    

</template>

<script>
    
    import axios from 'axios'
import data from 'css-tree/data'
    
    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/login/login/background.jpeg'),
                Username: "",
                Password: "",
                RealName:"",
                Identity:"",
                BeginYear:2018,
                EndYear:2021,
                leftValue:{},
                showLeftValue: false,
                opApi1 : {},
                opApi2 : {},
            }
        },
        methods:{
            getIdentity(){
                if ( this.Identity == "0" ) {
                    return "管理员"
                }
                else if ( this.Identity == "1" ) {
                    return "投资方"
                }
                else if ( this.Identity == "2" ) {
                    return "从业者"
                }
                else {
                    return "游客"
                }
            },
            Api1(){
                let that = this
            
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api1?BeginYear=2014&EndYear=2021"
                ).then(function(res){
                    console.log(res)
                    if (res.data.Code == 0){
                        let colors = ['#5470C6', '#EE6666'];
                        
                        let xData = [];
                        let yData1 = [];
                        let yData2 = [];
                        
                        for(let i = 0; i <= 2021 - 2014; i ++) {
                            xData.push(res.data.Data[i].Year)
                            yData1.push(res.data.Data[i].TotalAmount)
                            yData2.push(res.data.Data[i].TotalAmountPercentage)
                        }
                        console.log(">>>", xData, yData1, yData2)
                        that.opApi1 = {
                            color: colors,
                            tooltip: {
                                trigger: 'axis',
                                axisPointer: {
                                type: 'cross'
                                }
                            },
                            grid: {
                                right: '20%'
                            },
                            toolbox: {
                                feature: {
                                dataView: { show: true, readOnly: false },
                                restore: { show: true },
                                saveAsImage: { show: true }
                                }
                            },
                            legend: {
                                data: ['百分比', '市场份额']
                            },
                            xAxis: [
                                {
                                type: 'category',
                                axisTick: {
                                    alignWithLabel: true
                                },
                                // prettier-ignore
                                // data: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
                                    data: xData
                                }
                            ],
                            yAxis: [
                                {
                                type: 'value',
                                name: '市场份额',
                                position: 'right',
                                alignTicks: true,
                                offset: 80,
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
                                // data: [
                                //     2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3
                                // ]
                                data : yData1
                                },
                                {
                                name: '百分比',
                                type: 'line',
                                yAxisIndex: 1,
                                // data: [2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 6.2]
                                data: yData2
                                }
                            ]
                        };
                    }
                }).catch(function(err){
                    console.log(err)
                })
                
                
            },
            Api2(){
                let that = this
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api2?BeginYear=2010&EndYear=2021"
                ).then( function(res){
                    if (res.data.Code == 0){
                        that.opApi2 = {
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
                                smooth: true
                                }
                            ]
                        };
                    }
                }).catch(function(err){
                    console.log(err)
                })
            },
            Api3(){
                let that = this
                axios.post(
                    // TODO 范围查询
                    "http://127.0.0.1:1432/api/v1/api3?BeginYear=2018&EndYear=2021"
                ).then(function(res){
                    if (res.data.Code == 0) {
                        that.leftValue = res.data.Data
                        that.showLeftValue = true
                    }
                })
            }


        },
        mounted(){
            this.Username = this.$route.query.Username,
            this.RealName = this.$route.query.RealName,
            this.Identity = this.$route.query.Identity,
            this.Api1()
            this.Api2()
            this.Api3()
        }
    }

</script>


<style>
    .background{
        height:100%; 
        width:100%;
        top: 0%;
        left: 0%;
        z-index: -1;
        position: absolute;
    }

    .dataView{
        height: 100%;
        width: 100%;
        top: 0%;
        left: 0%;
        z-index: 1;
        position: absolute;
        display: flex;
        
        opacity: 0.7; 
        /* background-color: antiquewhite; */
    }    
        


        .content{
            height:100%; 
            width:80%;
            display: flex;
            position: absolute;     
            /* background-color: black;    */
        }

            
            .leftView{
                height:100%; 
                width:50%;
                left: 0%;
                display: flex;    
                position: absolute;
                /* background-color: red; */
            }
        
            .rightView{
                height:100%; 
                width:50%;
                left: 50%;
                display: flex;    
                position: absolute;     
                /* background-color: green;    */
            }

        .indexes{
            height:100%;
            width:20%;
            top: 0%;
            display: flex;
            position: absolute;
            left: 80%;
            background-color: rgb(124, 94, 16);
        }
.leftBox{
    top: 0%;
    height: 100%;
    width: 30%;
    position: absolute;

}
.rightBox{
    top: 0%;
    height: 100%;
    width: 70%;
    left: 30%;
    position: absolute;
}
.data0{
    top: 0%;
    height: 10%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: antiquewhite; */
}
.data1{
    height: 10%;
    top: 10%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: aqua; */
}
.data2{
    height: 10%;
    top: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: antiquewhite; */
}
.data3{
    height: 10%;
    top: 30%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: aqua; */
}
.data4{
    height: 60%;
    top: 40%;
    width: 100%;
    position: absolute;
    display: flex;
    /* background-color: antiquewhite; */
}

/* text */
.text{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    background-color: aquamarine;
}
.textIndex{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: left;  
    background-color: aquamarine;
}
.textTable{
    height:90%;
    width: 90%;
    top: 5%;
    left: 5%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    background-color: aquamarine;
}


.table1{
    top: 0%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: antiquewhite;
}
.table2{
    top: 20%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.table3{
    top: 40%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: antiquewhite;
}
.table4{
    top: 60%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: aqua;
}
.table5{
    top: 80%;
    height: 20%;
    width: 100%;
    position: absolute;
    display: flex;
    background-color: antiquewhite;
}
.tableLeft{
    height: 100%;
    width: 30%;
    left: 0%;
    position: absolute;
    display: flex;
}
.tableCenter{
    height: 100%;
    width: 40%;
    left: 30%;
    position: absolute;
    display: flex;
}
.tableRight{
    height: 100%;
    width: 30%;
    left: 70%;
    position: absolute;
    display: flex;
}

.top{
    height: 50%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    background-color: aliceblue;
}
.down{
    height: 50%;
    width: 100%;
    top:50%;
    position: absolute;
    display: flex;
    background-color: aliceblue;
}
 #input{
    text-align: left;
    height: 80%;
    width: 80%;
    top: 10%; 
    left: 0%;
    position: absolute;
 }
 
</style>