// To parse this JSON data, do
//
//     final amount = amountFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

Amount amountFromJson(String str) => Amount.fromJson(json.decode(str));

String amountToJson(Amount data) => json.encode(data.toJson());

class Amount {
  Amount({
    required this.code,
   required this.title,
   required this.status,
   required this.detail,
   required this.amountId,
   required this.net,
   required this.min,
   required this.max,
   required this.fee,
   required this.type,
   required this.logo,
   required this.sort,
   required this.category,
   required this.country,
   required this.currency,
   required this.protocol,
  });

  int code;
  String title;
  int status;
  String detail;
  int amountId;
  String net;
  int min;
  int max;
  int fee;
  String type;
  String logo;
  int sort;
  String category;
  String country;
  String currency;
  String protocol;

  factory Amount.fromJson(Map<String, dynamic> json) => Amount(
    code: json["code"],
    title: json["title"],
    status: json["status"],
    detail: json["detail"],
    amountId: json["amount_id"],
    net: json["net"],
    min: json["min"],
    max: json["max"],
    fee: json["fee"],
    type: json["type"],
    logo: json["logo"],
    sort: json["sort"],
    category: json["category"],
    country: json["country"],
    currency: json["currency"],
    protocol: json["protocol"],
  );

  Map<String, dynamic> toJson() => {
    "code": code,
    "title": title,
    "status": status,
    "detail": detail,
    "amount_id": amountId,
    "net": net,
    "min": min,
    "max": max,
    "fee": fee,
    "type": type,
    "logo": logo,
    "sort": sort,
    "category": category,
    "country": country,
    "currency": currency,
    "protocol": protocol,
  };
}
