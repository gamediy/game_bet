import 'package:bet/model/amount.dart';
import 'package:bet/utils/utils.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/utils.dart';

import '../../utils/http.dart';


class DepositLogic extends GetxController {

  TextEditingController amountTextController = new TextEditingController();
  List<dynamic> list=<dynamic>[].obs;
  List<double> money=<double>[10,50,100,200,300,500,1000,1500,2000,3000,5000,10000];
  var amount=0.00 ;
  var selectIndex=-1;
  var pressed=false;
  @override
  void onInit() {
    super.onInit();
      DioUtil.instance?.request("/game/amount_list",params: {
       "amount_type":"deposit"
     }).then((value){
        list.addAll(value["data"]);
        update();
      });






  }



  Deposit() async {

    if(amount<=0){
      return Utils.Toast("Deposit", "Wrong amount", false);
    }
    if(selectIndex<0){
      return Utils.Toast("Deposit", "Please select ", false);
    }

  }

}