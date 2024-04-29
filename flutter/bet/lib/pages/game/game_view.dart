import 'package:bet/pages/game/com/menu/menu_com_view.dart';
import 'package:bet/router/name.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:get/get.dart';
import 'package:getwidget/getwidget.dart';
import 'package:timer_count_down/timer_count_down.dart';

import '../../utils/values.dart';
import 'com/menu/menu_com_logic.dart';
import 'game_logic.dart';
import 'package:circular_countdown_timer/circular_countdown_timer.dart';
class GamePage extends StatefulWidget {
  const GamePage({Key? key}) : super(key: key);

  @override
  State<GamePage> createState() => _GamePage();
}
class _GamePage extends State<GamePage> {
  final logic = Get.put(GameLogic());
 final menu=Get.put(MenuComLogic());

  @override
  Widget build(BuildContext context) {
    return GetBuilder<GameLogic>(
      assignId: true,
      builder: (logic) {
        return Container(
          color: Values.Grey1,
          child: CustomScrollView(
            slivers: [
              SliverAppBar(
                  pinned: true,
                  title: Row(
                    children: [
                      Text(
                        "game@gmail.com",
                        style: TextStyle(fontSize: 14),
                      ),
                      SizedBox(
                        width: 10,
                      ),
                      Text(
                        "166",
                        style: TextStyle(fontSize: 18, color: Values.Amber),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      FaIcon(
                        FontAwesomeIcons.dollar,
                        size: 16,
                        color: Values.Amber,
                      )
                    ],
                  ),
                  actions: [
                    GFButton(
                      onPressed: () {},
                      child: Text(
                        "Withdraw",
                        style: TextStyle(color: Colors.white, fontSize: 16),
                      ),
                      color: Colors.black12,
                    ),
                    SizedBox(
                      width: 5,
                    ),
                    GFButton(
                      onPressed: () {},
                      child: Text(
                        "Deposit",
                        style: TextStyle(fontSize: 16),
                      ),
                      color: Colors.black12,
                    )
                  ]),
              SliverList(
                  delegate: SliverChildListDelegate([
                Container(
                  child: Placeholder(fallbackHeight: 200),
                ),
                SizedBox(
                  height: 5,
                ),
                Container(
                  margin: EdgeInsets.fromLTRB(10, 0, 10, 0),
                  decoration: BoxDecoration(
                    color: Colors.white,
                    //border: Border.all(width: 1, color: Colors.blue),
                    borderRadius: BorderRadius.all(
                      Radius.circular(5),
                    ),
                  ),
                  child: MenuCom(),
                ),
                Obx(() {
                  return     Container(
                    child: ListView.builder(
                      shrinkWrap: true,
                      itemCount: logic.list.length,
                      itemBuilder: (context, index) {
                        return Items(index);
                      },
                    ),
                  );
                })

              ]))
            ],
          ),
        );
      },
    );
  }
  Widget Items(index) {

   var cagte= menu.list.where((element) => element.id==menu.categoryId);
    var item=logic.list.value[index];
    var count=logic.count[index];
   var time=item.closeCountdown;
  if(item.status==2){
    time=item.openCountdown;
  }
    print(item.toJson());
    return InkWell(
      onTap: () {
        Get.toNamed(Name.Bet);
      },
      child: Container(
        margin: EdgeInsets.fromLTRB(10, 5, 10, 0),
        decoration: BoxDecoration(
            color: Values.White,
            borderRadius: BorderRadius.all(
              Radius.circular(5),
            )),
        child: new ListTile(
          leading: new CircleAvatar(
              foregroundColor: Colors.purple,
              backgroundColor: Colors.amber,
              child: Obx(() {
                return Text(
                 menu.list[menu.currentTypeIndex.value].name,
                  style: TextStyle(
                      color: Colors.white,
                      fontSize: 12,
                      fontWeight: FontWeight.bold),
                );
              })),
          title: new Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              new Text(
                item.name,
                style: new TextStyle(fontWeight: FontWeight.w100),
              ),
              new Text(
                item.statusStr,
                style: new TextStyle(color:item.status==1?Colors.green: Colors.red, fontSize: 14.0),
              ),
            ],
          ),
          subtitle: new Container(
            padding: const EdgeInsets.only(top: 5.0),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                new Text(
                  item.issue.toString(),
                  style: new TextStyle(color: Colors.grey, fontSize: 15.0),
                ),
                Countdown(
                  controller: count!,
                  seconds: time,
                  build: (BuildContext context, double time) =>
                      Text(time.toString()),
                  interval: Duration(milliseconds: 1000),
                  onFinished: () {
                    setState(() {
                      print("finished");
                      // print(logic.categoryId);
                      logic.restart(index);
                      count!.restart();


                    });

                  },
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }


}
