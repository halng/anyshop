//package com.app.anyshop.api.config;
//
//import java.util.HashMap;
//import java.util.Map;
//import org.springframework.http.HttpMethod;
//
//public class ExcludeUrlConfig {
//
//    private static final Map<String, HttpMethod> excludedUrl;
//
//    static {
//        excludedUrl = new HashMap<>();
//        excludedUrl.put("create-staff", HttpMethod.POST);
//        excludedUrl.put("login", HttpMethod.POST);
//        excludedUrl.put("image", HttpMethod.GET);
//        excludedUrl.put("active-account", HttpMethod.POST);
//        excludedUrl.put("inventory", HttpMethod.GET);
//    }
//
//    public static boolean isSecure(String requestPath, HttpMethod method) {
//        for (var ex : excludedUrl.entrySet()) {
//            if (requestPath.contains(ex.getKey()) && method.equals(ex.getValue())) {
//                return false;
//            }
//        }
//        return true;
//    }
//}