import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';
import 'package:getwidget/getwidget.dart';

import '../../../../utils/values.dart';
import 'menu_com_logic.dart';

class MenuCom extends StatefulWidget {

  MenuCom({Key? key}) : super(key: key);
  @override
  State<MenuCom> createState() => _MenuComComponentState();
}

class _MenuComComponentState extends State<MenuCom> with AutomaticKeepAliveClientMixin  {
 final logic = Get.put(MenuComLogic());

@override
void initState(){


}
  @override
  Widget build(BuildContext context) {


    return GetBuilder<MenuComLogic>(
      assignId: true,
      builder: (logic) {

      return  Container(
        width: double.infinity,
          child: Obx(() {
            return logic.body.value;
          }),
        );
      },
    );
  }
  //方法返回true
  @override
  bool get wantKeepAlive => true;

}


