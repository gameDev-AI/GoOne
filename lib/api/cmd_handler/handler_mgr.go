package cmd_handler

import (
	"errors"
	"sync"
)

/**
 * CmdHandlerMgr
 * @Description:
**/
type CmdHandlerMgr struct {
	httpHandlers map[string]WebHandlerFunc
	wsHandlers   map[string]WebHandlerFunc

	handler         RegCmdInterface
	handlersRWMutex sync.RWMutex
}

/**
* @Description: 创建cmdhandler管理器
* @return: *CmdHandler
* @Author: Iori
* @Date: 2022-08-11 17:54:43
**/
func NewCmdHandlerMgr() *CmdHandlerMgr {
	impl := &CmdHandlerMgr{}

	impl.httpHandlers = make(map[string]WebHandlerFunc)
	impl.wsHandlers = make(map[string]WebHandlerFunc)
	return impl
}

/**
* @Description: 初始化cmdhandler管理器
* @return: *CmdHandler
* @Author: Iori
* @Date: 2022-08-11 17:54:43
**/
func (self *CmdHandlerMgr) Init(impl RegCmdInterface) error {
	if impl == nil {
		return errors.New("RegCmd impl is nil!")
	}

	self.handler = impl
	return self.handler.Init()
}

/**
* @Description: reg http handler
* @param: key
* @param: value
* @Author: Iori
* @Date: 2022-04-26 17:31:33
**/
func (self *CmdHandlerMgr) HttpRegister(key string, value WebHandlerFunc) {
	self.handlersRWMutex.Lock()
	defer self.handlersRWMutex.Unlock()
	self.httpHandlers[key] = value

	return
}

/**
* @Description:  get http handlers
* @param: key
* @return: value
* @return: ok
* @Author: Iori
* @Date: 2022-04-26 17:31:21
**/
func (self *CmdHandlerMgr) GetHttpHandlers(key string) (value WebHandlerFunc, ok bool) {
	self.handlersRWMutex.RLock()
	defer self.handlersRWMutex.RUnlock()

	value, ok = self.httpHandlers[key]
	return
}

/**
* @Description: reg Ws Handlers
* @param: key
* @param: value
* @Author: Iori
* @Date: 2022-04-26 17:31:33
**/
func (self *CmdHandlerMgr) WsRegister(key string, value WebHandlerFunc) {
	self.handlersRWMutex.Lock()
	defer self.handlersRWMutex.Unlock()

	self.wsHandlers[key] = value
	return
}

/**
* @Description:  get ws handlers
* @param: key
* @return: value
* @return: ok
* @Author: Iori
* @Date: 2022-04-26 17:31:21
**/
func (self *CmdHandlerMgr) GetWsHandlers(key string) (value WebHandlerFunc, ok bool) {
	self.handlersRWMutex.RLock()
	defer self.handlersRWMutex.RUnlock()

	value, ok = self.wsHandlers[key]
	return
}
