class GameIssueRespone {
  int issue;
  DateTime openTime;
  DateTime closeTime;
  DateTime startTime;
  String openTimeStr;
  String closeTimeStr;
  String startTimeStr;
  int closeCountdown;
  int openCountdown;
  int status;
  String statusStr;
  DateTime timeNow;
  String date;
  String name;
  int gameCode;

  GameIssueRespone({
    required this.issue,
    required this.openTime,
    required this.closeTime,
    required this.startTime,
    required this.openTimeStr,
    required this.closeTimeStr,
    required this.startTimeStr,
    required this.closeCountdown,
    required this.openCountdown,
    required this.status,
    required this.statusStr,
    required this.timeNow,
    required this.date,
    required this.name,
    required this.gameCode,
  });

  factory GameIssueRespone.fromJson(Map<String, dynamic> json) {
    return GameIssueRespone(
      issue: json['issue'],
      openTime: DateTime.parse(json['open_time']),
      closeTime: DateTime.parse(json['close_time']),
      startTime: DateTime.parse(json['start_time']),
      openTimeStr: json['open_time_str'],
      closeTimeStr: json['close_time_str'],
      startTimeStr: json['start_time_str'],
      closeCountdown: json['close_countdown'],
      openCountdown: json['open_countdown'],
      status: json['status'],
      statusStr: json['status_str'],
      timeNow: DateTime.parse(json['time_now']),
      date: json['date'],
      name: json['name'],
      gameCode: json["game_code"]
    );
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['issue'] = this.issue;
    data['open_time'] = this.openTime.toIso8601String();
    data['close_time'] = this.closeTime.toIso8601String();
    data['start_time'] = this.startTime.toIso8601String();
    data['open_time_str'] = this.openTimeStr;
    data['close_time_str'] = this.closeTimeStr;
    data['start_time_str'] = this.startTimeStr;
    data['close_countdown'] = this.closeCountdown;
    data['open_countdown'] = this.openCountdown;
    data['status'] = this.status;
    data['status_str'] = this.statusStr;
    data['time_now'] = this.timeNow.toIso8601String();
    data['date'] = this.date;
    data['name'] = this.name;
    data["game_code"]=this.gameCode;
    return data;
  }
}
