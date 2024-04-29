import 'dart:html';
import 'package:bet/pages/bet/com/bet/bet_com_view.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:getwidget/getwidget.dart';
import 'package:bet/model/category.dart';
import 'package:bet/pages/bet/com/history/history_com_view.dart';
import 'package:bet/pages/bet/com/issue/issue_com_view.dart';


import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';

import '../../utils/values.dart';
import '../component/bottom.dart';
import 'bet_logic.dart';


class BetPage extends StatefulWidget {
  const BetPage({Key? key}) : super(key: key);

  @override
  State<BetPage> createState() => _BetPageState();
}

class _BetPageState extends State<BetPage> {
  final logic = Get.put(BetLogic());

  @override
  Widget build(BuildContext context) {
    ScreenUtil.init(context, designSize: const Size(390, 844));
    return GetBuilder<BetLogic>(
      assignId: true,
      builder: (logic) {
        return Scaffold(
            appBar: AppBar(
              title:Text("BTC/USDT(1)"),
              centerTitle: true,
              actions: [
                IconButton(onPressed: (){}, icon: FaIcon(FontAwesomeIcons.addressCard))
              ],
            ),
            body: Container(
              color: Values.Grey1,
              child: SingleChildScrollView(
                child: Column(
                  children: [
                    IssueCom(),
                    BetCom(),
                    HistoryCom(),
                  ],
                ),
              ),
            ));

      },
    );
  }
}
