const canvas=document.getElementById("ClockCanvas");
var ctx = canvas.getContext("2d");//建立上下文


canvas.height = 600;
canvas.width = 250;
ctx.scale(0.5,0.5);//初始化

var img_clockboard = new Image();
img_clockboard.src="src/ClockBoard.png"//初始化和绘制表盘

var img_Bolt = new Image();
img_Bolt.src="src/Bolt.png"//初始化插销

function First_Draw_Bolt(){//第一次绘制插销 主要是让图片加载完成后再显示  
    ctx.drawImage(img_Bolt,137.5,100);
}
setTimeout(First_Draw_Bolt,200)// NEED TO BE DEbug


var img_HourHand = new Image();//初始化分针和秒针
var img_MinuteHand = new Image();
img_HourHand.src = "src/HourHand.png";
img_MinuteHand.src = "src/MinuteHand.png";


//计算时间
var Now_Hour,Now_Minute,Now_Angle,Clock_Interval;
function get_hours(s) {
    return (s[0]-'0')*10+(s[1]-'0');
}
function get_minutes(s) {
    return (s[3]-'0')*10+(s[4]-'0');
}

function Draw_the_clock() {
    ctx.clearRect(0,800,500,500);
    ctx.drawImage(img_clockboard,0,800);
}

function Set_Time_to(Hours,Minutes) {
    Now_Hour=Hours;Now_Minute=Minutes;
    Now_Angle=Math.PI/6*(Hours+Minutes/60.0);
    Draw_the_clock();//初始化
    ctx.save();//显示时针
    ctx.translate(178,973);
    ctx.rotate(Math.PI/6*(Hours+Minutes/60.0));
    ctx.drawImage(img_HourHand,-16.5,-138);
    ctx.restore();    


    ctx.save();//显示分针
    ctx.translate(175.5,971);
    ctx.rotate(Math.PI/30*Minutes); 
    ctx.drawImage(img_MinuteHand,-3.5,-138);
    ctx.restore();
}



// ctx.globalCompositeOperation = "copy"
function SetTheClock() { //每一次根据时间调整钟表状态
    var __Time = getTime();
    var __hours = get_hours(__Time);
    var __minutes = get_minutes(__Time);
    __hours%=12;
    
    Set_Time_to(__hours,__minutes);
}
Clock_Interval=setInterval(SetTheClock,1);



var Bolt_x=137.5;
var Bolt_y=100;
var Bolt_Interval;
function Move_Bolt_A_Step(){
    ctx.clearRect(Bolt_x,Bolt_y,img_Bolt.width,img_Bolt.height);
    Bolt_y+=5;
    ctx.drawImage(img_Bolt,Bolt_x,Bolt_y);
}
function Move_the_Bolt() {
    if(Bolt_y>=700) {
        clearInterval(Bolt_Interval);
        Stop_The_Time();
    }
    else Move_Bolt_A_Step();
}
function Check_In_The_Bolt(e){ //判断是否点到了插销上面
    if(e.target!=canvas)return 0;
    if(e.offsetX>=68.75&&e.offsetX<=108.75&&e.offsetY<=150&&e.offsetY>=50)return 1;
    return 0;
}
function clistener(e){
    // console.log(e.offsetX,e.offsetY);
    if(Check_In_The_Bolt(e)) {
    //    console.log("nbl");
       Bolt_Interval=setInterval(Move_the_Bolt,1);
       removeEventListener("click",clistener);
    }
}
addEventListener("click",clistener);


function Check_In_The_Clock(e){ //判断是否点到了插销上面
    if(e.target!=canvas)return 0;
    if(e.offsetX>=0&&e.offsetX<=177.5&&e.offsetY<=577.5&&e.offsetY>=400)return 1;
    return 0;
}

function Stop_The_Time() {
    freezeTime();
    //clearInterval(Clock_Interval);
    addEventListener("click",olistener);
}

var Is_moving=false;
var Final_Angle;
var Hand_Interval;
function olistener(e) {
    if(Check_In_The_Clock(e)&&(!Is_moving))
    {
        
        var Leny=(488.75-e.offsetY),Lenx=(e.offsetX-88.75);
        var Len=Math.sqrt(Lenx*Lenx+Leny*Leny);
        var Cosx=Lenx/Len,Sinx=Leny/Len;
        // console.log("THE PLACE IS ",Lenx,Leny,Len,Cosx,Sinx);
        var Angle=Math.asin(Sinx);
        if      (Sinx>=0.0&&Cosx>=0.0)Final_Angle=Math.PI/2.0-Angle;
        else if (Sinx>=0.0&&Cosx<=0.0)Final_Angle=Math.PI*1.5+Angle;
        else if (Sinx<=0.0&&Cosx>=0.0)Final_Angle=Math.PI/2.0-Angle;
        else if (Sinx<=0.0&&Cosx<=0.0)Final_Angle=Math.PI*1.5+Angle;
        Is_moving=true;
        // console.log("I GOT",Sinx,Cosx,Angle,Final_Angle);
        meltTime();
        Hand_Interval=setInterval(Move_the_hand,0.1);
    }
}
function Move_the_hand() {
    if(Math.abs(Final_Angle-Now_Angle)<=0.05) {
        clearInterval(Hand_Interval);
        Is_moving=false;
        freezeTime();
    }
    // else console.log("NOW IT'S",Final_Angle,Now_Angle,Final_Angle-Now_Angle);
}