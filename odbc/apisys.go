package odbc

import "unsafe"
import "syscall"

var (
	mododbc32 = syscall.NewLazyDLL("odbc32.dll")

	procSQLAllocHandle     = mododbc32.NewProc("SQLAllocHandle")
	procSQLSetEnvAttr      = mododbc32.NewProc("SQLSetEnvAttr")
	procSQLDriverConnectW  = mododbc32.NewProc("SQLDriverConnectW")
	procSQLFreeHandle      = mododbc32.NewProc("SQLFreeHandle")
	procSQLDisconnect      = mododbc32.NewProc("SQLDisconnect")
	procSQLCancel          = mododbc32.NewProc("SQLCancel")
	procSQLExecDirectW     = mododbc32.NewProc("SQLExecDirectW")
	procSQLCloseCursor     = mododbc32.NewProc("SQLCloseCursor")
	procSQLFetch           = mododbc32.NewProc("SQLFetch")
	procSQLFetchScroll     = mododbc32.NewProc("SQLFetchScroll")
	procSQLSetStmtAttr     = mododbc32.NewProc("SQLSetStmtAttr")
	procSQLBindCol         = mododbc32.NewProc("SQLBindCol")
	procSQLSetConnectAttrW = mododbc32.NewProc("SQLSetConnectAttrW")
	procSQLEndTran         = mododbc32.NewProc("SQLEndTran")
	procSQLBindParameter   = mododbc32.NewProc("SQLBindParameter")
	procSQLMoreResults     = mododbc32.NewProc("SQLMoreResults")
	procSQLGetDescField    = mododbc32.NewProc("SQLGetDescField")
	procSQLGetDescRecW     = mododbc32.NewProc("SQLGetDescRecW")
	procSQLGetDiagRecW     = mododbc32.NewProc("SQLGetDiagRecW")
	procSQLColAttributeW   = mododbc32.NewProc("SQLColAttributeW")
	procSQLNumResultCols   = mododbc32.NewProc("SQLNumResultCols")
	procSQLGetData         = mododbc32.NewProc("SQLGetData")
	procSQLGetStmtAttr     = mododbc32.NewProc("SQLGetStmtAttr")
	procSQLSetDescFieldW   = mododbc32.NewProc("SQLSetDescFieldW")
)

func SQLAllocHandle(handleType SQLHandle, inputHandle syscall.Handle, outputHandle *syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLAllocHandle.Addr(), 3, uintptr(handleType), uintptr(inputHandle), uintptr(unsafe.Pointer(outputHandle)))
	ret = SQLReturn(r0)
	return
}

func SQLSetEnvAttr(environmentHandle syscall.Handle, attribute SQLINTEGER, valuePtr int32, stringLength SQLINTEGER) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLSetEnvAttr.Addr(), 4, uintptr(environmentHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLDriverConnect(connectionHandle syscall.Handle, windowHandle int, inConnString *uint16, inConnStringLength int16, outConnString *int, outConnStringLength int16, outConnStringPtr *int16, driverCompletion uint16) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall9(procSQLDriverConnectW.Addr(), 8, uintptr(connectionHandle), uintptr(windowHandle), uintptr(unsafe.Pointer(inConnString)), uintptr(inConnStringLength), uintptr(unsafe.Pointer(outConnString)), uintptr(outConnStringLength), uintptr(unsafe.Pointer(outConnStringPtr)), uintptr(driverCompletion), 0)
	ret = SQLReturn(r0)
	return
}

func SQLFreeHandle(handleType SQLHandle, handle syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLFreeHandle.Addr(), 2, uintptr(handleType), uintptr(handle), 0)
	ret = SQLReturn(r0)
	return
}

func SQLDisconnect(handle syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLDisconnect.Addr(), 1, uintptr(handle), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLCancel(statementHandle syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLCancel.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLExecDirect(statementHandle syscall.Handle, statementText *uint16, textLength int) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLExecDirectW.Addr(), 3, uintptr(statementHandle), uintptr(unsafe.Pointer(statementText)), uintptr(textLength))
	ret = SQLReturn(r0)
	return
}

func SQLCloseCursor(statementHandle syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLCloseCursor.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLFetch(statementHandle syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLFetch.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLFetchScroll(statementHandle syscall.Handle, fetchOrientation int16, fetchOffset int32) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLFetchScroll.Addr(), 3, uintptr(statementHandle), uintptr(fetchOrientation), uintptr(fetchOffset))
	ret = SQLReturn(r0)
	return
}

func SQLSetStmtAttr(statementHandle syscall.Handle, attribute int, valuePtr int32, stringLength int) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLSetStmtAttr.Addr(), 4, uintptr(statementHandle), uintptr(attribute), uintptr(valuePtr), uintptr(stringLength), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLBindCol(statementHandle syscall.Handle, columnNumber uint16, targetType CDataType, targetValuePtr uintptr, bufferLength SQLLEN, ind *SQLValueIndicator) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLBindCol.Addr(), 6, uintptr(statementHandle), uintptr(columnNumber), uintptr(targetType), uintptr(targetValuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(ind)))
	ret = SQLReturn(r0)
	return
}

func SQLSetConnectAttr(connectionHandle syscall.Handle, attribute SQLConnectionAttribute, valuePtr int32, bufferLength SQLINTEGER, stringLengthPtr *SQLINTEGER) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLSetConnectAttrW.Addr(), 5, uintptr(connectionHandle), uintptr(attribute), uintptr(valuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLengthPtr)), 0)
	ret = SQLReturn(r0)
	return
}

func SQLEndTran(handleType SQLHandle, handle syscall.Handle, completionType SQLTransactionOption) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLEndTran.Addr(), 3, uintptr(handleType), uintptr(handle), uintptr(completionType))
	ret = SQLReturn(r0)
	return
}

func SQLBindParameter(statementHandle syscall.Handle, parameterNumber SQLUSMALLINT, inputOutputType SQLBindParameterType, valueType CDataType, parameterType SQLDataType, columnSize SQLULEN, decimalDigits SQLSMALLINT, parameterValue SQLPOINTER, bufferLength SQLLEN, ind *SQLValueIndicator) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall12(procSQLBindParameter.Addr(), 10, uintptr(statementHandle), uintptr(parameterNumber), uintptr(inputOutputType), uintptr(valueType), uintptr(parameterType), uintptr(columnSize), uintptr(decimalDigits), uintptr(parameterValue), uintptr(bufferLength), uintptr(unsafe.Pointer(ind)), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLMoreResults(statementHandle syscall.Handle) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLMoreResults.Addr(), 1, uintptr(statementHandle), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLGetDescField(descriptorHandle syscall.Handle, recNumber int16, fieldIdentifier int16, valuePtr uintptr, bufferLength int, lengthPtr *int) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLGetDescField.Addr(), 6, uintptr(descriptorHandle), uintptr(recNumber), uintptr(fieldIdentifier), uintptr(valuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(lengthPtr)))
	ret = SQLReturn(r0)
	return
}

func SQLGetDescRec(descriptorHandle syscall.Handle, recNumber int16, name *uint16, bufferLength int16, stringLengthPtr *int16, typePtr *int16, subTypePtr *int16, lengthPtr *int, precisionPtr *int16, scalePtr *int16, nullablePtr *int16) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall12(procSQLGetDescRecW.Addr(), 11, uintptr(descriptorHandle), uintptr(recNumber), uintptr(unsafe.Pointer(name)), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLengthPtr)), uintptr(unsafe.Pointer(typePtr)), uintptr(unsafe.Pointer(subTypePtr)), uintptr(unsafe.Pointer(lengthPtr)), uintptr(unsafe.Pointer(precisionPtr)), uintptr(unsafe.Pointer(scalePtr)), uintptr(unsafe.Pointer(nullablePtr)), 0)
	ret = SQLReturn(r0)
	return
}

func SQLGetDiagRec(handleType SQLHandle, inputHandle syscall.Handle, recNumber int16, sqlState uintptr, nativeErrorPtr *int, messageText uintptr, bufferLength int16, textLengthPtr *int16) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall9(procSQLGetDiagRecW.Addr(), 8, uintptr(handleType), uintptr(inputHandle), uintptr(recNumber), uintptr(sqlState), uintptr(unsafe.Pointer(nativeErrorPtr)), uintptr(messageText), uintptr(bufferLength), uintptr(unsafe.Pointer(textLengthPtr)), 0)
	ret = SQLReturn(r0)
	return
}

func SQLColAttribute(statementHandle syscall.Handle, columnNumber SQLUSMALLINT, fieldIdentifier SQLColAttributeType, characterAttribute uintptr, bufferLength SQLSMALLINT, stringLengthPtr *SQLSMALLINT, numericAttributePtr *SQLLEN) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall9(procSQLColAttributeW.Addr(), 7, uintptr(statementHandle), uintptr(columnNumber), uintptr(fieldIdentifier), uintptr(characterAttribute), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLengthPtr)), uintptr(unsafe.Pointer(numericAttributePtr)), 0, 0)
	ret = SQLReturn(r0)
	return
}

func SQLNumResultCols(statementHandle syscall.Handle, columnCount *int16) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall(procSQLNumResultCols.Addr(), 2, uintptr(statementHandle), uintptr(unsafe.Pointer(columnCount)), 0)
	ret = SQLReturn(r0)
	return
}

func SQLGetData(statementHandle syscall.Handle, colNum uint16, targetType CDataType, targetValuePtr uintptr, bufferLength SQLLEN, ind *SQLValueIndicator) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLGetData.Addr(), 6, uintptr(statementHandle), uintptr(colNum), uintptr(targetType), uintptr(targetValuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(ind)))
	ret = SQLReturn(r0)
	return
}

func SQLGetStmtAttr(statementHandle syscall.Handle, attribute SQLStatementAttribute, valuePtr uintptr, bufferLength SQLINTEGER, stringLengthPtr *SQLINTEGER) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLGetStmtAttr.Addr(), 5, uintptr(statementHandle), uintptr(attribute), uintptr(valuePtr), uintptr(bufferLength), uintptr(unsafe.Pointer(stringLengthPtr)), 0)
	ret = SQLReturn(r0)
	return
}

func SQLSetDescField(descriptorHandle syscall.Handle, recNum SQLSMALLINT, fieldIdentifier SQLSMALLINT, valuePtr int32, bufferLength SQLINTEGER) (ret SQLReturn) {
	r0, _, _ := syscall.Syscall6(procSQLSetDescFieldW.Addr(), 5, uintptr(descriptorHandle), uintptr(recNum), uintptr(fieldIdentifier), uintptr(valuePtr), uintptr(bufferLength), 0)
	ret = SQLReturn(r0)
	return
}
