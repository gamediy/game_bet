import 'dart:core';

import 'package:bet/pages/game/com/menu/menu_com_logic.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';
import 'package:timer_count_down/timer_controller.dart';
import 'package:timer_count_down/timer_count_down.dart';

import '../../model/game_issue.dart';
import '../../router/name.dart';
import '../../utils/http.dart';
import '../../utils/values.dart';
import 'com/menu/menu_com_view.dart';

class GameLogic extends GetxController with GetTickerProviderStateMixin {
  late TabController groupTabController;
  late TabController typeTabController;
  var count = <CountdownController>[];
  CountdownController? controller;
  var int = 0.obs;
  var list = <GameIssueRespone>[].obs;
  final menu = Get.put(MenuComLogic());
  var categoryId = 0;

  @override
  void onInit() {
    controller = new CountdownController(autoStart: true);
    // TODO: implement onInit
    groupTabController = TabController(length: 2, vsync: this);
    typeTabController = TabController(length: 4, vsync: this);

    ever(menu.categoryId, (_) => getGameList(menu.categoryId.value));
    getGameList(1001);
    super.onInit();
  }

  getGameList(categoryid) async {

    count = <CountdownController>[];
    categoryId = categoryid;
    print("gamelist");
    var res = await DioUtil.instance?.request("/game/game_list",
        method: DioMethod.post, data: {"category": categoryid});
    if (res["code"] == 500) {
      return;
    }
    list.value = [];
    for (var item in res["data"]) {
      var i = GameIssueRespone.fromJson(item);
      count.add(new CountdownController(autoStart: true));
      list.value.add(i);
    }
  }

  restart(index) async{
    var item=list[index];
    var res=await DioUtil.instance?.request("/game/issue",method: DioMethod.post,data: {
      "game_code":item.gameCode
    });

    list[index]=GameIssueRespone.fromJson(res["data"]);
  }
}
