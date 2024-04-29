import 'dart:html';
import 'dart:math';

import 'package:bet/pages/deposit/deposit_logic.dart';
import 'package:bet/router/name.dart';
import 'package:bet/utils/http.dart';
import 'package:bet/utils/utils.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:get/get.dart';
import 'package:qr_flutter/qr_flutter.dart';

class DepositPage extends StatefulWidget {
  @override
  _DepositPage createState() => new _DepositPage();
}

class _DepositPage extends State<DepositPage> {
  void ShowBs(int selectIndex) async {

    try{
      var item=logic.list[selectIndex];
      int min=item["min"];
      int max=item["max"];
      logic.amount=double.parse(logic.amountTextController.text);
      print(logic.amount);
      if(logic.amount<min||logic.amount>max){

        Utils.Toast("Deposit", 'Deposit amount minimum:$min maximum:$max', false);
        return;

      }
     var res=await DioUtil.instance?.request("/deposit/deposit",data: {
        "amount":logic.amount,
        "amount_item_code":item["code"]
      });
      int code=res["code"];
      if (code!=200){
        Utils.Toast("Deposit", res["message"], false);
      }
      var data=res["data"];
      Get.bottomSheet(Container(
        color: Colors.grey[50],
        height: 400,
        padding: EdgeInsets.all(10),
        child: Stack(

          children: [
            Expanded(
                flex: 1,
                child: Container(
                    height: 20,
                    margin: EdgeInsets.only(top: 30),
                    child: Row(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      mainAxisAlignment: MainAxisAlignment.spaceAround,
                      children: [
                        Text(
                          "Address",
                          style: Theme.of(context).textTheme.headline1,

                        ),
                        Text(
                          data["address"],
                          style: Theme.of(context).textTheme.headline2,
                        ),
                        InkWell(
                          onTap: (){

                          },
                          child: Text("Copy",
                              style: TextStyle(
                                  color: Colors.blueAccent, fontSize: 13,fontWeight: FontWeight.bold)),
                        ),
                      ],
                    ))),
            Positioned(
                top: 80,
                left:   0.5.sw-150*0.5,

              child: Container(
                  height: 150,
                  width: 150,
                  margin: EdgeInsets.only(top: 0),
                  child:QrImage(
                    data:  data["address"],
                    version: QrVersions.auto,
                    size: 150.0,
                  ),
            )),

            Align(
              alignment: Alignment.bottomLeft,
                child: Container(

                  padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
                  
                  height: 120,

                  width: double.infinity,
                  alignment: Alignment.bottomRight,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        "Tips",
                        style: TextStyle(color: Colors.black,fontWeight: FontWeight.bold),
                      ),
                      RichText(
                        text:TextSpan(
                          style: TextStyle(color: Colors.grey),
                          text: "test testseteste steste test testseteste steste test testseteste stestetest testseteste stestetest testseteste steste"
                        )
                      )
                    ],
                  ),
                ))
            // Positioned(
            //   bottom: 5,
            //   left: 10,
            //   right: 10,
            //   child: MaterialButton(
            //     height: 60,
            //     minWidth: 320,
            //     child: Text(
            //       "Config",
            //       style: Theme.of(context).textTheme.bodyText2,
            //     ),
            //     color: Colors.purple,
            //     onPressed: () {},
            //   ),
          ],
        ),
      ));
    }
    catch(e){
     print(e);
    }


  }
  @override
  void initState() {
    super.initState();
  }
  final logic = Get.put(DepositLogic());
  @override
  Widget build(BuildContext context) {
    ScreenUtil.init(context, designSize: const Size(390, 844));
    // TODO: implement build
    return Scaffold(
        appBar: AppBar(
          centerTitle: true,
          primary: false,
          title: Text("Deposit"),
          leading: BackButton(),
            actions:[
             IconButton(onPressed: (){
               Get.toNamed(Name.DepositRecord);
             },icon: Icon(Icons.receipt_outlined),)
            ]
        ),
        body: GetBuilder<DepositLogic>(

          assignId: true,
          builder: (logic) {
            return Container(
              padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
              color:  Color(0xf5f7fb),
              child: Column(
                children: [
                  Container(
                    color: Colors.white,
                      height: 60.0,
                      padding: EdgeInsets.fromLTRB(15, 10, 15, 10),
                      child: TextFormField(
                        controller: logic.amountTextController,
                        decoration: new InputDecoration(
                          hintText: 'Select amount or enter ',
                        ),
                        validator: (money) {
                          if (money == null) {
                            return "Please enter amount";
                          }
                          var m = double.parse(money);
                          if (m <= 0) {
                            return "Please enter amount";
                          }
                        },
                        onSaved: (value) {
                          print("222");
                          setState(() {
                            logic.amount = double.parse(value!);
                          });

                        },
                      )),
                  Container(
                    color: Colors.white,
                    child: GridView.builder(

                        padding: EdgeInsets.fromLTRB(15, 10, 15, 10),
                        itemCount: logic.money.length,
                        shrinkWrap: true,
                        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 4,
                          mainAxisSpacing: 12,
                          crossAxisSpacing: 20,
                          childAspectRatio: 1 / 0.5,
                        ),
                        itemBuilder: (context, index) {
                          return Material(
                            color: Colors.white,
                            child: InkWell(
                              onTap: () {
                                logic.amountTextController.text =
                                    logic.money[index].toString();
                                setState(() {
                                  logic.selectIndex = index;
                                });
                              },
                              child: Container(

                                decoration: logic.selectIndex == index
                                    ? BoxDecoration(
                                        color: Colors.purple,
                                        borderRadius: BorderRadius.all(
                                          Radius.circular(0),
                                        ))
                                    : BoxDecoration(
                                  color: Colors.grey[500]
                                ),
                                width: 60,
                                height: 30,
                                alignment: Alignment.center,
                                child: Text(
                                  logic.money[index].toString(),
                                  style: TextStyle(color: Colors.white),
                                ),
                              ),
                            ),
                          );
                        }),
                  ),
                  Container(
                    margin: EdgeInsets.only(top: 10),
                    color: Colors.white,
                    child: ListView.builder(
                        shrinkWrap: true,
                        itemCount: logic.list.length,
                        itemBuilder: (context, index) {
                          return Column(
                            children: [
                              ListTile(
                                title: Text(logic.list[index]["title"]),
                                leading: Image.network(
                                    "https://img0.baidu.com/it/u=2799820758,1545153940&fm=253&fmt=auto&app=138&f=JPEG?w=540&h=336"),
                                trailing: Icon(Icons.chevron_right),
                                onTap: () {
                                  ShowBs(index);
                                },
                              ),
                              if (index % 2 == 0) Divider()
                            ],
                          );
                        }),
                  ),
                ],
              ),
            );
          },
        ));
  }
}
