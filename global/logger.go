package global

import "my-service/pkg/logger"

var (
	Logger        *logger.Logger
	DaoLogger     *logger.Logger
	ServiceLogger *logger.Logger
	ModelLogger   *logger.Logger
	ApiLogger     *logger.Logger
)
