import 'package:bet/pages/deposit/record/record_logic.dart';
import 'package:easy_refresh/easy_refresh.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class DepositRecordPage extends StatelessWidget {
   DepositRecordPage({Key? key}) : super(key: key);
  final logic = Get.put(DepositRecordLogic());
  @override
  Widget build(BuildContext context) {
    return GetBuilder<DepositRecordLogic>(
      builder: (logic) {
        return Scaffold(
          appBar: AppBar(
            centerTitle: true,
            primary: false,
            automaticallyImplyLeading : true,
            title: Text("Deposit record"),
            leading: IconButton(
              onPressed: (){

                print("back");
                Get.back();
              },
              icon: new Icon(Icons.arrow_back_ios),
            ),
          ),

          body: Container(
            child: EasyRefresh(
                onRefresh: logic.refreshData,
                onLoad: logic.loadMoreData,
                child: ListView.builder(
                  itemCount: logic.dataList.length,
                  itemBuilder: (context, index) {
                    return new ListTile(
                      leading: new CircleAvatar(
                        foregroundColor: Theme.of(context).primaryColor,
                        backgroundColor: Colors.teal,
                        child: Text("USDT",style: TextStyle(color: Colors.white,fontSize: 12),),
                      ),
                      title: new Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: <Widget>[
                          new Text(
                            logic.dataList[index]["amount"].toString(),
                            style: new TextStyle(fontWeight: FontWeight.bold),
                          ),
                          new Text(
                            logic.dataList[index]["status_text"],
                            style: new TextStyle(
                                color:logic.GetStatusColor(logic.dataList[index]["status"]), fontSize: 14.0),
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
                              logic.dataList[index]["title"],
                              style:
                              new TextStyle(color: Colors.grey, fontSize: 15.0),
                            ),
                            new Text(

                              logic.dataList[index]["created_at"]??"",
                              style:
                              new TextStyle(color: Colors.grey, fontSize: 15.0),
                            ),
                          ],
                        ),
                      ),
                    );
                  },
                )),
          ),
        );
      },
    );
  }
}





