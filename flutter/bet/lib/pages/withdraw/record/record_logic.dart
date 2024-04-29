import 'package:bet/utils/http.dart';
import 'package:dio/dio.dart';
import 'package:easy_refresh/easy_refresh.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class WithdrawRecordLogic extends GetxController {

  List dataList=[];
  var pageSize = 20;
  var pageIndex = 1;
  bool hadMore = true;
  EasyRefreshController easyRefreshController = EasyRefreshController(
    controlFinishRefresh: true,
    controlFinishLoad: true,
  );

  @override
  void onInit() async {
    getList();
    super.onInit();
  }
  MaterialColor GetStatusColor(int status){
    if (status==1){
    return Colors.amber;
    }
    else if(status==2){
    return Colors.red;
    }

    return Colors.green;
  }
  void getList() async {
    var res = await DioUtil.instance?.request("/game/deposit_record", data: {
      "PageSize": this.pageSize,
      "PageIndex": this.pageIndex,
      "Status": 1
    });
    List list = res["data"];
    if (list.length < pageSize) {
      hadMore = false;
    } else {
      hadMore = true;
    }
    if (dataList.isEmpty) {
      easyRefreshController.finishRefresh();
    } else {
      easyRefreshController.finishLoad(
          hadMore ? IndicatorResult.success : IndicatorResult.noMore);
    }
    dataList.addAll(list);
    update();
  }

  Future refreshData() async{
    pageIndex=1;
    dataList=[];
    getList();
  }
  Future loadMoreData() async{
    if (hadMore) {
      pageIndex++;
      getList();
    } else {
      easyRefreshController.finishLoad(IndicatorResult.noMore);
    }
  }
  @override
  void onClose() {
    easyRefreshController.dispose();
    super.onClose();
  }
}
