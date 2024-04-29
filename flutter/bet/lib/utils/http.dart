import 'package:bet/router/name.dart';
import 'package:bet/utils/utils.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:get/get.dart' as getx;
/// 请求方法
enum DioMethod {
  get,
  post,
  put,
  delete,
  patch,
  head,
}

class AjaxResult {
  int? code;
  String? message;
  Object? data;

  AjaxResult({this.code, this.message, this.data});

  factory AjaxResult.fromJson(Map<String, dynamic> json) {
    return AjaxResult(
      code: json['code'],
      message: json['message'],
      data: json['data'],
    );
  }
}

class DioUtil {
  /// 单例模式
  static   DioUtil? _instance;

  factory DioUtil() => _instance ?? DioUtil._internal();

  static DioUtil get instance => _instance ?? DioUtil._internal();

  /// 连接超时时间
  static const int connectTimeout = 60 * 1000;

  /// 响应超时时间
  static const int receiveTimeout = 60 * 1000;

  /// Dio实例
  static Dio _dio = Dio();

  /// 初始化
  DioUtil._internal() {
    // 初始化基本选项
    BaseOptions options = BaseOptions(
        baseUrl: 'http://127.0.0.1:7001/app/',
        connectTimeout: connectTimeout,
        receiveTimeout: receiveTimeout);
    _instance = this;
    // 初始化dio
    _dio = Dio(options);
    // 添加拦截器
    _dio.interceptors.add(InterceptorsWrapper(
        onRequest: _onRequest, onResponse: _onResponse, onError: _onError));
  }

  /// 请求拦截器
  Future<void> _onRequest(RequestOptions options, RequestInterceptorHandler handler) async {
    //EasyLoading.show(status: 'loading...');

    if (options.path.contains("/user/")) {
      options.baseUrl = "http://127.0.0.1:8081/api";
    }
    else if(options.path.contains("/order/")){
      options.baseUrl = "http://127.0.0.1:8082/api";
    }
    else if(options.path.contains("/game/")){
      options.baseUrl = "http://127.0.0.1:8083/api";
    }
    else if(options.path.contains("/deposit/")){
      options.baseUrl = "http://127.0.0.1:8084/api";
    }
    print(options.path);

    // // 对非open的接口的请求参数全部增加userId
    // if (!options.path.contains("open")) {
    //   options.queryParameters["userId"] = "xxx";
    // }
    // 头部添加token
    var prefs = await SharedPreferences.getInstance();
    options.headers["Authorization"] = "Bearer "+prefs.getString(Utils.TokenKey).toString();
    // 更多业务需求
    handler.next(options);
    // super.onRequest(options, handler);
  }

  /// 相应拦截器
  void _onResponse(Response response, ResponseInterceptorHandler handler) async {
    // 请求成功是对数据做基本处理
    if (response.statusCode == 200) {
      // ....
    } else if (response.statusCode == 401) {
      print("401");
      // ....
    }
   // EasyLoading.dismiss();
    handler.next(response);
  }

  /// 错误处理
  void _onError(DioError error, ErrorInterceptorHandler handler) {
    EasyLoading.dismiss();
    handler.next(error);
  }

  /// 请求类
  Future<T> request<T>(
    String path, {
    DioMethod method = DioMethod.post,
    Map<String, dynamic>? params,
    data,
        showLoading=false,
    CancelToken? cancelToken,
    Options? options,
    ProgressCallback? onSendProgress,
    ProgressCallback? onReceiveProgress,
  }) async {
    const _methodValues = {
      DioMethod.get: 'get',
      DioMethod.post: 'post',
      DioMethod.put: 'put',
      DioMethod.delete: 'delete',
      DioMethod.patch: 'patch',
      DioMethod.head: 'head'
    };
    options ??= Options(method: _methodValues[method]);
    try {
      Response response;
      response = await _dio.request(path,
          data: data,
          queryParameters: params,
          cancelToken: cancelToken,
          options: options,
          onSendProgress: onSendProgress,
          onReceiveProgress: onReceiveProgress);

      return response.data;
    } on DioError catch (e) {

      if(e.response?.statusCode==401||e.response?.statusCode==403){
        if(path.contains("/user/login")){
          return e.response?.data;
        }
        getx.Get.offAllNamed(Name.Login);
        dynamic data= {
          "code":401

        };
        return  Future(() => data);
      }
      Utils.Toast("Error", e.message,false);
      return e.response?.data;
    }
  }

  /// 开启日志打印
  /// 需要打印日志的接口在接口请求前 DioUtil.instance?.openLog();
  void openLog() {
    _dio.interceptors
        .add(LogInterceptor(responseHeader: false, responseBody: true));
  }
}
