// To parse this JSON data, do
//
//     final category = categoryFromJson(jsonString);

import 'package:meta/meta.dart';
import 'dart:convert';

Category categoryFromJson(Map<String, dynamic>  str) => Category.fromJson(str);

String categoryToJson(Category data) => json.encode(data.toJson());

class Category {
  Category({
    required this.id,
    required this.logo,
    required this.name,
    required this.status,
    required this.parentId,
    required this.sort,
  });

  int id;
  String logo;
  String name;
  int status;
  int parentId;
  int sort;

  factory Category.fromJson(Map<String, dynamic> json) => Category(
    id: json["id"],
    logo: json["logo"],
    name: json["name"],
    status: json["status"],
    parentId: json["parent_id"],
    sort: json["sort"],
  );

  Map<String, dynamic> toJson() => {
    "id": id,
    "logo": logo,
    "name": name,
    "status": status,
    "parent_id": parentId,
    "sort": sort,
  };
}
