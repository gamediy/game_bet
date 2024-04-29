import 'package:bet/pages/bet/com/issue/issue_com_logic.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';
import 'package:indexed_list_view/indexed_list_view.dart';

import 'package:timer_count_down/timer_count_down.dart';

import '../../../../utils/values.dart';

class IssueCom extends StatefulWidget {
  const IssueCom({Key? key}) : super(key: key);

  @override
  State<IssueCom> createState() => _IssueComState();
}

class _IssueComState extends State<IssueCom> {
  @override
  Widget build(BuildContext context) {
    var logic = Get.put(IssueComLogic());
    return GetBuilder<IssueComLogic>(
      assignId: true,
      builder: (logic) {
        return Container(
          child: Column(
            children: [
              SizedBox(
                height: 5,
              ),
              Container(
                  width: double.infinity,
                  height: 80,
                  child: Card(
                    child: Container(
                      child: SingleChildScrollView(
                        reverse: true,
                        scrollDirection: Axis.horizontal,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: logic.issueList,
                        ),
                      ),
                    ),
                  )),

              Card(
                  child: Column(
                    children: [
                      Container(
                        height: 40,
                      padding: EdgeInsets.only(left: 5),
                        decoration: BoxDecoration(
                            border: Border(
                                bottom: BorderSide(
                                  color: Values.Grey1,
                                ))),
                        alignment: Alignment.centerLeft,
                        child: Text(
                          "Bet issue",
                          style: TextStyle(
                            fontWeight: FontWeight.w500,
                            fontSize: 16,
                            color: Values.Grey,
                          ),
                        ),
                      ),
                      Container(
                          height: 60,
                          alignment: Alignment.center,
                          padding: EdgeInsets.fromLTRB(5, 10, 5, 10),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceBetween,
                            children: [

                              Text(
                                logic.issue,
                                style: TextStyle(
                                    color: Colors.amber,
                                    fontSize: 20,
                                    fontWeight: FontWeight.bold),
                              ),
                              Countdown(
                                controller: logic.controller!,
                                seconds:logic.time,
                                build: (BuildContext context, double time) =>
                                    Container(
                                        height: 30,
                                        width: 50,
                                        alignment: Alignment.center,
                                        decoration: BoxDecoration(
                                            color: Colors.redAccent,
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(5))),
                                        child: Text(
                                          time.toString(),
                                          style: TextStyle(
                                            color: Colors.white,
                                            fontSize: 18,
                                            fontWeight: FontWeight.bold,
                                          ),
                                        )),
                                interval: Duration(milliseconds: 1000),
                                onFinished: () {
                                  setState(() {
                                    logic.GetIssue();
                                    logic.controller!.restart();
                                  });
                                },
                              ),
                              Text(
                                logic.status,
                                style: TextStyle(
                                  fontWeight:  FontWeight.bold,
                                    color: logic.status=="Betting"?Colors.green:Colors.redAccent, fontSize: 18
                                ),
                              )
                            ],
                          )),
                    ],
                  )),
            ],
          ),
        );
      },
    );
  }
}
