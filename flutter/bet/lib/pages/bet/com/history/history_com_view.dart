import 'package:flutter/material.dart';
import 'package:get/get.dart';
import '../../../../utils/values.dart';
import 'history_com_logic.dart';

class HistoryCom extends StatelessWidget {
  final logic = Get.put(HistoryComLogic());

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Container(
        child: Column(
          children: [
            Container(
              height: 40,
              alignment: Alignment.centerLeft,
              decoration: BoxDecoration(
                border: Border(
                    bottom: BorderSide(
                        color: Values.Grey1,width: 1
                    )
                )
              ),
              child: Text(" Bet history",style: TextStyle(
                  fontWeight: FontWeight.w500,
                  fontSize: 16,
                  color: Values.Grey,
              ),),
            ),

            Container(

              child: ListView.builder(
                  shrinkWrap: true,
                  itemCount: 5,
                  itemBuilder: (context,index){
                    return new ListTile(

                      shape: Border(
                          bottom: BorderSide(
                              color: Values.Grey1,width: 1
                          )
                      ),


                      leading: new CircleAvatar(
                        foregroundColor: Theme.of(context).primaryColor,
                        backgroundColor: Values.RedBet,
                        child: Text("1023",style: TextStyle(color: Colors.white,fontSize: 13),),
                      ),
                      title: new Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: <Widget>[
                          new Text(
                            "Big",
                            style: new TextStyle(fontWeight: FontWeight.bold),
                          ),
                          new Text(
                            "wait",
                            style: new TextStyle(
                                color:Colors.grey, fontSize: 14.0),
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
                              "98",
                              style:
                              new TextStyle(color: Colors.grey, fontSize: 15.0),
                            ),
                            new Text(

                              "2022-02-15 12:23:23",
                              style:
                              new TextStyle(color: Colors.grey, fontSize: 15.0),
                            ),
                          ],
                        ),
                      ),
                    );
                  }),
            )
          ],
        ),
      ),
    );
  }
}
