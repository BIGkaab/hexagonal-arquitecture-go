package impl

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/dummy"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testMapper = NewTaskMapperImpl()

func TestTaskDtoToDomain_OK(t *testing.T) {
	task := testMapper.TaskDtoToDomain(dummy.TaskDto)
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskDomain, task)
}

func TestTaskDomainToDto_OK(t *testing.T) {
	task := testMapper.TaskDomainToDto(dummy.TaskDomain)
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskDto, task)
}

func TestTaskDomainToEntity(t *testing.T) {
	task := testMapper.TaskDomainToEntity(dummy.TaskDomain)
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskEntity, task)
}

func TestTaskEntityToDomain_OK(t *testing.T) {
	task := testMapper.TaskEntityToDomain(dummy.TaskEntity)
	assert.NotNil(t, task)
	assert.Equal(t, dummy.TaskDomain, task)
}
