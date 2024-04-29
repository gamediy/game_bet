import 'package:bet/model/category.dart';
import 'package:bet/utils/http.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import '../../../../utils/values.dart';
class MenuComLogic extends GetxController with GetTickerProviderStateMixin {

  final List<Tab> typeList = <Tab>[];
  var currentTypeIndex = 0.obs;
  var  categoryId=0.obs;
  late   List<Tab> betList = <Tab>[];
  late var body=Column().obs;
  late TabController typeController;
  List<Category> list=<Category>[];

  List<Tab> GetTypeList() {
    var ty= list.where((element) => element.parentId==0);
    for (var value in ty) {
      typeList.add(Tab(
        text: value.name,
      ));
    }
    return typeList;
  }


  void init(){

  }
  @override
  void onInit()  {
    print("menu");

body.value= Column(children: [
  Center(
      child: CircularProgressIndicator())
],);


    typeController = TabController(vsync: this, length:0);
     DioUtil.instance?.request("/game/game_category",method: DioMethod.post).then((res){

       for(var item in res["data"]){
         var i=Category.fromJson(item);
         list.add(i);
       }
       GetTypeList();
       typeController = TabController(vsync: this, length: typeList.length);
      getTab();

     });

    super.onInit();
  }

  @override
  void onClose() {
    typeController.dispose();
    super.onClose();
  }

  void getTab(){
    body.value=Column(children: [
      Align(
        alignment: Alignment.centerLeft,
        child: Container(
          width: 1.sw,
          child: TabBar(
            isScrollable: true,
            controller: typeController,
            tabs: typeList,
            labelColor: Colors.blue,
            indicator: UnderlineTabIndicator(
                borderSide: BorderSide(
                  width: 4.0,
                  color: Values.MainColor,
                )),
            onTap: (index) {
              print(index);
              print("select");
              currentTypeIndex.value = index;
              categoryId.value=list[index].id;

            },
          ),
        ),
      ),
      Container(height:1,color: Values.Grey1)
    ],);


  }

}
