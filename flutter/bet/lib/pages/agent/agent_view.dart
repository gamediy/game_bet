import 'package:bet/pages/component/bottom.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

import 'agent_logic.dart';

class AgentPage extends StatelessWidget {
  final logic = Get.put(AgentLogic());

  @override
  Widget build(BuildContext context) {
    return Container(
       child: Text("agent"),
     );


  }
}
