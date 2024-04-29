import 'package:bet/utils/http.dart';
import 'package:bet/utils/values.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';
import 'package:indexed_list_view/indexed_list_view.dart';

import 'package:timer_count_down/timer_controller.dart';
import 'package:timer_count_down/timer_count_down.dart';

class IssueComLogic extends GetxController with GetTickerProviderStateMixin {
  var gameCode = 10000;
  List<Widget> issueList = [];
  int countdown = 0;
  String issue = "";
  String status = "";
  Widget Issue = CircularProgressIndicator();
  GlobalKey dataKey = GlobalKey();
  int time = 999;
  String select = "";
   CountdownController controller= new CountdownController(autoStart: true);

  @override
  void onInit() {
    GetIssue();
    issueList = [
      Padding(
          padding: EdgeInsets.fromLTRB(0, 0, 0.5.sw - 5, 0),
          child: CircularProgressIndicator())
    ];
    OpenList();

    WidgetsBinding.instance?.addPostFrameCallback((callback) {});

    super.onInit();
  }

  Future<List<Widget>?> OpenList() async {
    var res = await DioUtil.instance?.request("/game/game_open", data: {
      "game_code": gameCode,
      "limit": 20,
      "index": 1,
      "status": 9,
    });
    var list = res["data"];
    issueList = [];
    for (var item in list) {
      var i = int.tryParse(item["open_result"]);
      issueList.add(OpenItem(item["issue"].toString(), i!));
    }
    update();
  }

  Widget OpenItem(issue, int openresult) {
    return Row(
      key: issue == select ? dataKey : null,
      children: [
        ClipOval(
          child: Container(
            width: 60,
            height: 60,
            color: openresult > 4 ? Colors.red : Colors.blue,
            child: Stack(
              children: [
                Align(
                  alignment: Alignment.center,
                  child: Text(
                    issue,
                    style: TextStyle(
                      color: Colors.white,
                      fontWeight: FontWeight.w500,
                      fontSize: 20,
                    ),
                  ),
                ),
                Align(
                  alignment: Alignment.bottomCenter,
                  child: Text(
                    openresult.toString(),
                    style: TextStyle(
                      fontSize: 13,
                      color: Values.Grey,
                    ),
                  ),
                )
              ],
            ),
          ),
        ),
        SizedBox(
          width: 5,
        )
      ],
    );
  }

  void GetIssue() async {
    var res = await DioUtil.instance
        .request("/game/issue", method: DioMethod.post, data: {
      "game_code": gameCode,
    });
    var item = res["data"];
    countdown = item["close_countdown"];
    issue = item["issue"].toString();
    status = item["status_str"].toString();
    int open_countdown = item["open_countdown"];
    time=countdown;
    if (item["status"] == 2) {
      time=open_countdown;
    }

    update();
  }
}
