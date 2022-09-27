<template>
    <div class="AllDataView">
        <div class="background">
            <img :src="imgSrcBK" height="100%" width="100%"/>
        </div>

        <div class="dataView">
        
            <div class="content">
            
                <div class="leftView">
                    <e-charts :option="opApi4"/>
                </div>
                <div class="centerView">
                    <div class="top1">
                        <e-charts :option="opApi1"/>
                    </div>
                    <div class="down1">
                        <div id="scroll_in2" class="htp-list_scroll_in" style="height: 100%;overflow-y:hidden;font-size:20px;">
                            
                            <table>
                                <tr>
                                    <td colspan="3">
                                        <h1>
                                            2021年-2022年预制菜行业融资情况
                                        </h1>
                                    </td>
                                </tr>
                                <tr>
                                    <td style="height: 20%;"> <h3>品牌</h3> </td>
                                    <td> <h3>时间</h3> </td>
                                    <td> <h3> 轮次</h3> </td>
                                    <td> <h3>投资商</h3> </td>
                                </tr>
                                <tr v-for="data in Api5Value">
                                    <td style="height: 20%;">
                                        {{data.Brand}}
                                    </td>
                                    <td>
                                        {{data.Time}}
                                    </td>
                                    <td>
                                        {{data.FundingRound}}
                                    </td>
                                    <td>
                                        {{data.Investor}}
                                    </td>
                                </tr>
                            </table>

                        </div>
                    </div>
                </div>
                <div class="rightView">
                    <div class="top0">
                        
                        <div class="table0Box1">
                            <div class="text0">
                                <h3>
                                    欢迎{{getIdentity()}} {{RealName}}
                                </h3>
                            </div>
                        </div>
                        <div class="table0Box2">
                            <div class="text0">
                                Id : {{Id}}
                                </div>
                        </div>
                        <div class="table0Box3">
                            <div class="text0">
                                用户名 : {{Username}}
                            </div>
                        </div>
                    </div>

                    <div class="top2">
                        <div class="table1" style="font-size:20px;">
                            <div class="table1Box" >
                                <div class="text" style="text-align: left;">2018年-2021年餐饮就业人数行业平均报酬</div>
                            </div>
                            <div class="table1Box">
                                <div class="table1BoxData1"><div class="text">时间</div></div>
                                <div class="table1BoxData2"><div class="text">平均报酬</div></div>
                                <div class="table1BoxData3"><div class="text">就业人数</div></div>
                            </div>
                            <div v-for="data in Api3Value" class="table1Box">
                                <div class="table1BoxData1"><div class="text">{{data.Year}}</div></div>
                                <div class="table1BoxData2"><div class="text">{{data.AvgSalary}}</div></div>
                                <div class="table1BoxData3"><div class="text">{{data.CountSalary}}</div></div>
                            </div>
                        </div>
                        
                    </div>
                    <div class="down2">    
                        <e-charts :option="opApi2"/>
                    </div>
                </div>
                    
            </div>


            <div class="indexes">
                <div class="text">
                    <h3>

                        火<br />
                        锅<br />
                        数<br />
                        据<br />
                        大<br />
                        屏<br />
                        页
                        
                        <div class="text">
                            <button @click="logout">注销</button>
                        </div>
                    </h3>
                </div>
                
            </div>
        
        </div>
    </div>
    

</template>

<script>
    
    import axios from 'axios'
    
    export default{
        name: "login",

        data(){
            return{
                // 背景
                imgSrcBK:require('../../assets/background.jpeg'),
                rightValue:{},
                showLeftValue: false,
                Api3Value:{},
                Api5Value:{},
                opApi1 : {},
                opApi2 : {},
                opApi4 : {},
                Username: "",
                RealName:"",
                Identity:"",
                Id:""
            }
        },
        methods:{
            getIdentity(){
                if (this.Identity == '0') { return "管理员"}
                if (this.Identity == '1') { return "投资方"}
                if (this.Identity == '2') { return "从业者"}
                if (this.Identity == '3') { return "游客"}
            },
            logout(){
                this.$cookies.set("camp-session", "", -1)
                
                this.$router.push({
                    path: `/`,
                    query: {
                        msg: "注销成功"
                    },
                })
                
            },
            Api1(){
                let that = this
            
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api1?BeginYear=2014&EndYear=2021"
                ).then(function(res){
                    
                    if (res.data.Code == 0){
                        let colors = ['#559955', '#FF9900'];
                        
                        let xData = [];
                        let yData1 = [];
                        let yData2 = [];
                        
                        for(let i = 0; i <= 2021 - 2014; i ++) {
                            xData.push(res.data.Data[i].Year)
                            yData1.push(res.data.Data[i].TotalAmount)
                            yData2.push(res.data.Data[i].TotalAmountPercentage)
                        }
                        
                        that.opApi1 = {
                            title : {
                                text:"2014年-2021年中国餐饮市场规模"
                            },
                            color: colors,
                            tooltip: {
                                trigger: 'axis',
                                axisPointer: {
                                type: 'cross'
                                }
                            },
                            grid: {
                                right: '25%'
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
                                name: '市场份额(万亿)',
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
                                name: '市场份额(万亿)',
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
                        that.Api3Value = res.data.Data
                        that.showLeftValue = true
                    }
                })
            },
            Api4(){
                let that = this
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api4"
                ).then( function(res) {
                    if (res.data.Code == 0) {
                        that.opApi4 = {
                            title: {
                                text: '中国火锅品牌梯队分布',
                            },
                            series: [
                                {
                                type: 'funnel',
                                min: 0,
                                max: 100,
                                minSize: '0%',
                                maxSize: '100%',
                                sort: 'ascending',
                                gap: 2,
                                label: {
                                    show: true,
                                    position: 'inside',
                                    color: '#555555',
                                    fontSize:12.5,
                                },
                                labelLine: {
                                    length: 10,
                                    lineStyle: {
                                    width: 1,
                                    type: 'solid'
                                    }
                                },
                                itemStyle: {
                                    borderColor: '#FF9900',
                                    borderWidth: 5
                                },
                                emphasis: {
                                    label: {
                                    fontSize: 20
                                    }
                                },
                                data: [
                                    { 
                                        value: 20, 
                                        name: res.data.Data[0],
                                        itemStyle: {
                                            color: "#FF9900",
                                            borderColor: "none",
                                        },
                                    },
                                    { 
                                        value: 40, 
                                        name: res.data.Data[1],
                                        itemStyle: {
                                            color: "#FF7700",
                                            borderColor: "none",
                                        },
                                    },
                                    { 
                                        value: 60,
                                        name: res.data.Data[2],
                                        itemStyle: {
                                            color: "#FF5500",
                                            borderColor: "none",
                                        },
                                    },
                                    {
                                        value: 80,
                                        name: res.data.Data[3],
                                        itemStyle: {
                                            color: "#FF3300",
                                            borderColor: "none",
                                        },
                                     }
                                ]
                                }
                            ]
                            };
                    }
                }).catch(function(err){
                    console.log(err)
                })
            },
            Api5(){
                let that = this
                axios.post(
                    "http://127.0.0.1:1432/api/v1/api5"
                ).then(function(res){
                    if ( res.data.Code == 0 ) {
                        that.Api5Value = res.data.Data
                    }
                }).catch(function(err){
                    console.log(err)
                })
            },    // 添加自动滚动
            /* 
            Id   需要滚动的区域 id名称
            */
            autoSroll(Id) {
            // flag 为true时停止滚动
                var flag = false;
                // 定时器
                var timer;
                function roll() {
                    var h = -1;
                    timer = setInterval(function() {
                        flag = true;
                        if (flag) {
                            var current = document.getElementById(Id).scrollTop;
                            if (current == h) {
                                //滚动到底端,返回顶端
                                h = 0;
                                document.getElementById(Id).scrollTop = h + 1;
                            } else {
                                //以25ms/3.5px的速度滚动
                                h = current;
                                document.getElementById(Id).scrollTop = h + 1;
                            }
                            
                        }
                        //获取当前滚动条高度
                        }, 50);
                    }
                // setTimeout(function() {
                    //滚动区域内单击鼠标 滚动暂停 再次点击鼠标 继续滚动
                    document.getElementById(Id).onclick = () => {
                    
                    if (flag) {
                        flag = false;
                        clearInterval(timer);
                    } else {
                        roll();
                    }
                    };
                    roll();
                // }, 2000);
                },    
        },
        mounted(){
            this.Id = this.$route.query.Id 
            this.Api1()
            this.Api2()
            this.Api3()
            this.Api4()
            this.Api5()
            this.autoSroll("scroll_in2")
            let that = this
            axios.post(
                "http://127.0.0.1:1432/api/v1/whoAmI?",{                        
                    Id : that.Id
                }
            ).then(function(res){
                if (res.data.Code == 0 ){
                    that.Username = res.data.Data.Username
                    that.RealName = res.data.Data.RealName
                    that.Identity = res.data.Data.Identity
                } else {
                    
                }
            }).catch(function(err){
                console.log(err)
            })

        },
    }

</script>


<style>

 .AllDataView{
    height:100%; 
    width:100%;
    top: 0%;
    left: 0%;
    position: absolute;
 }   
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
        
        opacity: 0.750; 
        /* background-color: antiquewhite; */
    }    

        .content{
            height:100%; 
            width:95%;
            display: flex;
            position: absolute;     
            /* background-color: black;    */
        }

            .leftView{
                height:100%; 
                width:25%;
                left: 0%;
                display: flex;    
                position: absolute;
                background-color: rgb(255, 255, 255);
            }
            .centerView{
                height:100%; 
                width:45%;
                left: 25%;
                display: flex;    
                position: absolute;
                background-color: aliceblue;
            }      
            .rightView{
                height:100%; 
                width:30%;
                left: 70%;
                display: flex;    
                position: absolute;        
                background-color: azure;
            }

        .indexes{
            height:100%;
            width:5%;
            top: 0%;
            display: flex;
            position: absolute;
            left: 95%;
            background-color: rgb(124, 94, 16);
        }


/* text */
.text{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: center;  
    overflow-y:hidden
}
.text0{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 25px; 
    text-align: center;  
    overflow-y:hidden
}
.textIndex{
    height:100%;
    width: 100%;
    position: absolute;
    line-height: 100px; 
    text-align: left;  
}
.top0{
    height: 15%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    flex-direction: row;
    background-color: rgb(240,240,240)
}
.top1{
    height: 50%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
    background-color: rgb(250,250,250)
}
.top2{
    height: 35%;
    width: 100%;
    top:15%;
    position: absolute;
    display: flex;
    background-color:rgb(245, 245, 245);
}
.down1{
    height: 50%;
    width: 100%;
    top:50%;
    position: absolute;
    display: flex;
    background-color:rgb(245, 245, 245);
}
.down2{
    height: 50%;
    width: 100%;
    top:50%;
    position: absolute;
    display: flex;
    background-color: rgb(250,250,250);
}

.htp-list_scroll_in{
    height: 100%;
    width: 100%;
    top:0%;
    position: absolute;
    display: flex;
}

.table1{
    height: 100%;
    width: 100%;
    position: absolute;
    top: 0%;
    left: 0%;
    display: flex;
    flex-direction:column;
}
.table1Box{
    width: 100%;
    height: 15%;
    top: auto;
    display: flex;
    flex-direction: column;
}
.table0Box1{
    width: 100%;
    height: 50%;
    top:0%;
    display: flex;
    position: absolute;
}
.table0Box2{
    width: 100%;
    height: 25%;
    top: 50%;
    position: absolute;
    display: flex;
}
.table0Box3{
    height: 25%;
    width: 100%;
    top: 75%;
    position: absolute;
    display: flex;
}
.table1BoxData1{
    width: 30%;
    height: 100%;
    left: 0%;
    position: absolute;
    display: flex;
}
.table1BoxData2{
    width: 40%;
    height: 100%;
    left: 30%;
    position: absolute;
    display: flex;
}
.table1BoxData3{
    width: 30%;
    height: 100%;
    left: 70%;
    position: absolute;
    display: flex;
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

