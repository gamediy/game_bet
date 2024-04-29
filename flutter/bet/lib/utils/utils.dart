import 'package:flutter/material.dart';
import 'package:get/get.dart';

class Utils{
  static final String TokenKey="Token";

  static Toast(String title,String content,bool success,[Function(SnackbarStatus)? snackbarStatus=null]){
    Get.snackbar(title,content ,
      colorText: Colors.white,
      backgroundColor: Colors.purple,
      borderRadius: 10,
      icon:Icon(success?Icons.done_rounded:Icons.error,color: Colors.white),
      margin: EdgeInsets.all(10),

        snackbarStatus:(status){
        if(snackbarStatus!=null){
          snackbarStatus(status!);
        }
        },
      backgroundGradient: LinearGradient(colors: [Colors.deepPurple,Colors.purple]),
      borderColor: Colors.purple,
      borderWidth: 1,
      boxShadows: [
        BoxShadow(
          color: Colors.purple,
        ),
      ],

      //forwardAnimationCurve: Curves.bounceInOut,
      shouldIconPulse: false,

     overlayBlur: 1,
      /*overlayColor: Colors.grey,*/
      padding: EdgeInsets.all(20),


    );
  }



}