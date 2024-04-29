import 'package:bet/model/category.dart';
import 'package:bet/pages/bet/model/BetMenu.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

import '../../utils/ws.dart';

class BetLogic extends GetxController with GetTickerProviderStateMixin {
  var category=0.obs;
  var gameCoode=1000;

  @override
  void onInit() async {
   var ws= WebSocketUtility.Instance().initWebSocket(onOpen: (){
      print("OPEN WS");
      WebSocketUtility.Instance().initHeartBeat();
    }, onMessage: (data){
      print(data);
    }, onError: (){});
    ever(category, (value)  {
      print(value);
      print("Category");
    });
   WebSocketUtility.Instance().initHeartBeat();
    super.onInit();


    @override
    void onClose() {
      super.onClose();
    }
  }
}
