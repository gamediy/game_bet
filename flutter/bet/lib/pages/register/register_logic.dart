import 'package:bet/utils/utils.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/utils.dart';

import '../../utils/http.dart';
class RegisterLogic extends GetxController {



  var email="";
  var password="";
  var password2="";
  @override
  void onInit() {

    super.onInit();
  }
  register() async{
    var data={
      "email": this.email,
    "password": this.password,
    "password2": this.password2
    };
    var res=await DioUtil.instance?.request("/user/register",data: data);
    int code = res["code"];
    if(code==200){
      Utils.Toast("Regisetr", "Success going sing in",true,(status){
        if(status==SnackbarStatus.CLOSED){
          Get.offAllNamed("/login");
        }
      });
    }
    else{
      Utils.Toast("Regisetr", res["message"],false);
    }




  }

}
