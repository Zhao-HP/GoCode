package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

var param = map[string]string{
	"fromApp":            "storage-app",
	"sourceHost":         "test-deployment-b8cd66679-c8bfg 110.110.110.110",
	"title":              "java.lang.NullPointerException 空指针异常",
	"exceptionClassName": "java.lang.NullPointerException ",
	"temacode":           "yzo2o",
	"content": "java.lang.NullPointerException\n" +
		"at com.ziroom.crm.cm.service.crm.zwhite.impl.ZwhiteServiceImpl.zWhiteLoanUnwind(ZwhiteServiceImpl.java:272)\n" +
		"at com.ziroom.crm.cm.service.crm.zwhite.impl.ZwhiteServiceImpl.notifyContractToZwhite(ZwhiteServiceImpl.java:184)\n" +
		"at sun.reflect.GeneratedMethodAccessor844.invoke(Unknown Source)\n" +
		"at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n" +
		"at java.lang.reflect.Method.invoke(Method.java:498)\n" +
		"at org.springframework.aop.support.AopUtils.invokeJoinpointUsingReflection(AopUtils.java:317)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.invokeJoinpoint(ReflectiveMethodInvocation.java:190)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:157)\n" +
		"at org.springframework.transaction.interceptor.TransactionInterceptor$1.proceedWithInvocation(TransactionInterceptor.java:99)\n" +
		"at org.springframework.transaction.interceptor.TransactionAspectSupport.invokeWithinTransaction(TransactionAspectSupport.java:281)\n" +
		"at org.springframework.transaction.interceptor.TransactionInterceptor.invoke(TransactionInterceptor.java:96)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:179)\n" +
		"at com.alibaba.druid.support.spring.stat.DruidStatInterceptor.invoke(DruidStatInterceptor.java:72)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:179)\n" +
		"at org.springframework.aop.interceptor.ExposeInvocationInterceptor.invoke(ExposeInvocationInterceptor.java:92)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:179)\n" +
		"at org.springframework.aop.framework.JdkDynamicAopProxy.invoke(JdkDynamicAopProxy.java:207)\n" +
		"at com.sun.proxy.$Proxy253.notifyContractToZwhite(Unknown Source)\n" +
		"at com.ziroom.crm.cm.task.contract.CloseZwhiteContractTask.notifyContractToZwhite(CloseZwhiteContractTask.java:27)\n" +
		"at com.ziroom.crm.cm.task.contract.CloseZwhiteContractTask$$FastClassBySpringCGLIB$$296f082e.invoke(<generated>)\n" +
		"at org.springframework.cglib.proxy.MethodProxy.invoke(MethodProxy.java:204)\n" +
		"at org.springframework.aop.framework.CglibAopProxy$CglibMethodInvocation.invokeJoinpoint(CglibAopProxy.java:717)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:157)\n" +
		"at org.springframework.aop.aspectj.MethodInvocationProceedingJoinPoint.proceed(MethodInvocationProceedingJoinPoint.java:85)\n" +
		"at com.ziroom.crm.cm.interceptor.TaskAspect.aroundMethod(TaskAspect.java:75)\n" +
		"at sun.reflect.GeneratedMethodAccessor660.invoke(Unknown Source)\n" +
		"at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n" +
		"at java.lang.reflect.Method.invoke(Method.java:498)\n" +
		"at org.springframework.aop.aspectj.AbstractAspectJAdvice.invokeAdviceMethodWithGivenArgs(AbstractAspectJAdvice.java:621)\n" +
		"at org.springframework.aop.aspectj.AbstractAspectJAdvice.invokeAdviceMethod(AbstractAspectJAdvice.java:610)\n" +
		"at org.springframework.aop.aspectj.AspectJAroundAdvice.invoke(AspectJAroundAdvice.java:68)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:179)\n" +
		"at com.alibaba.druid.support.spring.stat.DruidStatInterceptor.invoke(DruidStatInterceptor.java:72)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:179)\n" +
		"at org.springframework.aop.interceptor.ExposeInvocationInterceptor.invoke(ExposeInvocationInterceptor.java:92)\n" +
		"at org.springframework.aop.framework.ReflectiveMethodInvocation.proceed(ReflectiveMethodInvocation.java:179)\n" +
		"at org.springframework.aop.framework.CglibAopProxy$DynamicAdvisedInterceptor.intercept(CglibAopProxy.java:653)\n" +
		"at com.ziroom.crm.cm.task.contract.CloseZwhiteContractTask$$EnhancerBySpringCGLIB$$92f4d298.notifyContractToZwhite(<generated>)\n" +
		"at sun.reflect.GeneratedMethodAccessor783.invoke(Unknown Source)\n" +
		"at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)\n" +
		"at java.lang.reflect.Method.invoke(Method.java:498)\n" +
		"at org.springframework.scheduling.support.ScheduledMethodRunnable.run(ScheduledMethodRunnable.java:65)\n" +
		"at org.springframework.scheduling.support.DelegatingErrorHandlingRunnable.run(DelegatingErrorHandlingRunnable.java:54)\n" +
		"at org.springframework.scheduling.concurrent.ReschedulingRunnable.run(ReschedulingRunnable.java:81)\n" +
		"at java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:511)\n" +
		"at java.util.concurrent.FutureTask.run(FutureTask.java:266)\n" +
		"at java.util.concurrent.ScheduledThreadPoolExecutor$ScheduledFutureTask.access$201(ScheduledThreadPoolExecutor.java:180)\n" +
		"at java.util.concurrent.ScheduledThreadPoolExecutor$ScheduledFutureTask.run(ScheduledThreadPoolExecutor.java:293)\n" +
		"at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1149)\n" +
		"at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:624)\n" +
		"at java.lang.Thread.run(Thread.java:748)\n",
}

const (
	ip      = "10.30.105.120"
	port    = "8083"
	apiPath = "api/v2/newTrigger.json"
)

func joinUrl() string {
	return fmt.Sprintf("http://%s:%s/%s", ip, port, apiPath)
	//return "http://" + ip + ":" + port + apiPath
}

func TestLogLocal(t *testing.T) {
	for i := 0; i < 22; i++ {
		sendSingleRequest()
	}
}

func sendSingleRequest() {
	url := joinUrl()
	fmt.Println(url)
	b, _ := json.Marshal(param)
	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("发送请求异常: ", err)
		return
	}
	fmt.Println(resp.Body)
}
