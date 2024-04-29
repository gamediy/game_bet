import 'package:bet/pages/bet/bet_logic.dart';
import 'package:bet/utils/http.dart';
import 'package:bet/utils/ws.dart';
import 'package:flutter/material.dart';

import 'package:get/get.dart';

import '../../../../utils/values.dart';

class BetComLogic extends GetxController {
  String play = "";
  TextEditingController inputController=TextEditingController();
  List<double> moneyList = [
    10,
    50,
    100,
    300,
    500,
    1000,
    3000,
    5000,
    0.1,
    0.25,
    0.5,
    1
  ];
  int moneyIndex = -1;
  double balance=0;
  @override
  void onInit() async {
    WebSocketUtility.Instance().onMessage!((data){
      print(data);
    });
    var res=await DioUtil.instance.request("/user/userinfo");
    balance=res["balance"];
    update();
  }
  void SelectPlay(play) {
    this.play = play;
  }
   SelectMoneyIndex(index) {
    this.moneyIndex = index;
    var money=moneyList[index];
    inputController.text=money.toString();
    if(money<=1){
      inputController.text=(money*balance).toString();
    }
    update();

  }
}
