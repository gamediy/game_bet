import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:get/get.dart';
import 'package:simple_animations/simple_animations.dart';

import '../../../../utils/values.dart';
import 'bet_com_logic.dart';

class BetCom extends StatelessWidget {
  final logic = Get.put(BetComLogic());

  @override
  Widget build(BuildContext context) {
    return GetBuilder<BetComLogic>(
      assignId: true,
      builder: (logic) {
        return Column(
          children: [
            Card(
              child: Column(
                children: [
                  Container(
                    height: 40,
                    margin: EdgeInsets.only(left: 5),
                    decoration: BoxDecoration(
                        border: Border(
                            bottom: BorderSide(
                      color: Values.Grey1,
                    ))),
                    alignment: Alignment.centerLeft,
                    child: Text(
                      "Bet click",
                      style: TextStyle(
                        fontWeight: FontWeight.w500,
                        fontSize: 16,
                        color: Values.Grey,
                      ),
                    ),
                  ),
                  LoopAnimationBuilder<double>(
                      tween: Tween(begin: 20.0, end: 30.0), // 100.0 to 200.0
                      duration: const Duration(seconds: 1), // for 1 second
                      builder: (context, value, _) {
                        return Container(
                          margin: EdgeInsets.fromLTRB(0, 5, 5, 0),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [
                              SizedBox(
                                height: 10,
                              ),
                              InkWell(
                                onTap: () {
                                  logic.SelectPlay("Big");
                                },
                                child: Container(
                                  decoration: BoxDecoration(
                                      color: Color.fromRGBO(232, 65, 24, 1),
                                      boxShadow: [
                                        BoxShadow(
                                            color: Colors.red,
                                            blurRadius: value)
                                      ],
                                      border: Border.all(
                                          color: Color.fromRGBO(194, 54, 22, 1),
                                          width: logic.play == "Big" ? 3 : 0)),
                                  height: 120,
                                  width: 0.5.sw - 12,
                                  child: Stack(
                                    children: [
                                      Align(
                                        child: Container(
                                            alignment: Alignment.center,
                                            height: 40,
                                            width: 0.5.sw - 12,
                                            color:
                                                Color.fromRGBO(194, 54, 22, 1),
                                            child: Text(
                                              "BIG",
                                              style: Values.FontWhite16w,
                                              textAlign: TextAlign.center,
                                            )),
                                        alignment: Alignment.topLeft,
                                      ),
                                      Align(
                                        child: Container(
                                            child: Text(
                                          "1.98%",
                                          style: Values.FontWhite18w,
                                        )),
                                        alignment: Alignment.center,
                                      ),
                                      Align(
                                        child: Text(
                                          "1800",
                                          style: Values.FontWhite16w,
                                        ),
                                        alignment: Alignment.bottomCenter,
                                      ),
                                      if (logic.play == "Big")
                                        Positioned(
                                            right: 5,
                                            bottom: 5,
                                            child: FaIcon(
                                              FontAwesomeIcons.check,
                                              color: Colors.white,
                                              size: 32,
                                            ))
                                    ],
                                  ),
                                ),
                              ),
                              InkWell(
                                onTap: () {
                                  logic.SelectPlay("Small");
                                },
                                child: Container(
                                  decoration: BoxDecoration(
                                      color: Values.BlueBet,
                                      boxShadow: [
                                        BoxShadow(
                                            color: Values.BlueBet,
                                            blurRadius: value)
                                      ],
                                      border: Border.all(
                                          color: Values.BlueBet2,
                                          width:
                                              logic.play == "Small" ? 3 : 0)),
                                  height: 120,
                                  width: 0.5.sw - 12,
                                  child: Stack(
                                    children: [
                                      Align(
                                        child: Container(
                                            alignment: Alignment.center,
                                            height: 40,
                                            width: 0.5.sw - 12,
                                            color: Values.BlueBet2,
                                            child: Text(
                                              "Small",
                                              style: Values.FontWhite16w,
                                              textAlign: TextAlign.center,
                                            )),
                                        alignment: Alignment.topLeft,
                                      ),
                                      Align(
                                        child: Text(
                                          "1.98%",
                                          style: Values.FontWhite18w,
                                        ),
                                        alignment: Alignment.center,
                                      ),
                                      Align(
                                        child: Text(
                                          "1800",
                                          style: Values.FontWhite16w,
                                        ),
                                        alignment: Alignment.bottomCenter,
                                      ),
                                      if (logic.play == "Small")
                                        Positioned(
                                            right: 5,
                                            bottom: 5,
                                            child: FaIcon(
                                              FontAwesomeIcons.check,
                                              color: Colors.white,
                                              size: 32,
                                            ))
                                    ],
                                  ),
                                ),
                              ),
                            ],
                          ),
                        );
                      }),
                  Container(
                    child: GridView.builder(
                        padding: EdgeInsets.fromLTRB(5, 5, 5, 0),
                        itemCount: 12,
                        shrinkWrap: true,
                        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 4,
                          mainAxisSpacing: 5,
                          crossAxisSpacing: 5,
                          childAspectRatio: 1 / 0.5,
                        ),
                        itemBuilder: (context, index) {
                          return Material(
                            child: InkWell(
                              onTap: () {
                              logic.SelectMoneyIndex(index);
                              },
                              child:Stack(
                                children: [
                                  Expanded(
                                    child: Container(
                                      decoration:
                                      BoxDecoration(
                                          color: logic.moneyIndex==index?Values.MainColor:Values.Grey
                                      ),
                                      width: 150,
                                      height: 50,
                                      alignment: Alignment.center,
                                      child: Text(
                                        logic.moneyList[index]<=1? (logic.moneyList[index]*100).toString()+"%":  logic.moneyList[index].toString(),
                                        style: TextStyle(color: Colors.white),
                                      ),
                                    ),
                                  ),
                                  if(index==logic.moneyIndex)
                                    Positioned(
                                        right: 5,
                                        bottom: 5,
                                        child: FaIcon(
                                          FontAwesomeIcons.check,
                                          color: Colors.white,
                                          size: 16,
                                        ))



                                ],
                              ),
                            ),
                          );
                        }),
                  ),
                  Container(
                    padding: EdgeInsets.fromLTRB(5, 5, 5, 5),
                    height: 60,
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Container(
                          height: 50,
                          width: 0.6.sw,
                          child: TextFormField(
                            controller: logic.inputController,
                            decoration: new InputDecoration(
                                labelStyle: new TextStyle(
                                    fontSize: 15.0,
                                    color: Color.fromARGB(255, 93, 93, 93)),
                                border: OutlineInputBorder(),
                                suffixIcon: new IconButton(
                                  icon: new Icon(Icons.delete),
                                  onPressed: () {
                                    logic.inputController.text="";
                                  },
                                )),
                          ),
                        ),
                        Container(
                          height: 55.h,
                          width: 0.35.sw,
                          child: MaterialButton(
                            onPressed: () {},
                            color: Values.MainColor,
                            child: Text(
                              "Bet confirm",
                              style: TextStyle(
                                  color: Colors.white,
                                  fontSize: 16,
                                  fontWeight: FontWeight.w500),
                            ),
                            shape: new RoundedRectangleBorder(
                                borderRadius: new BorderRadius.circular(10.0)),
                          ),
                        )
                      ],
                    ),
                  ),
                  Container(
                    margin: EdgeInsets.fromLTRB(5,0, 5, 5),
                    height: 30,
                    child: Row(
                      children: [
                        Text("Balance",style: TextStyle(
                          color: Values.Grey,
                          fontSize: 13,
                        ),),
                        SizedBox(width: 5,),
                        Text(logic.balance.toString(),style: TextStyle(
                          color: Colors.red,
                          fontSize: 13,
                          fontWeight: FontWeight.bold
                        ),)

                      ],
                    ),
                  )
                ],
              ),
            ),
          ],
        );
      },
    );
  }
}
