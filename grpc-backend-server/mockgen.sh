#!/bin/bash
# example
# mockgen -source domain/repository/#{repository}.go -destination mock/mock_repository/#{repository}_mock.go

mockgen -source domain/repository/contract_manager_repository.go -destination mock/mock_repository/contract_manager_repository_mock.go
mockgen -source domain/repository/harp_sys_sequence_repository.go -destination mock/mock_repository/harp_sys_sequence_repository_mock.go
mockgen -source domain/repository/item_units_repository.go -destination mock/mock_repository/item_units_repository_mock.go
mockgen -source domain/repository/order_details_repository.go -destination mock/mock_repository/order_details_repository_mock.go
mockgen -source domain/repository/order_items_repository.go -destination mock/mock_repository/order_items_repository_mock.go
mockgen -source domain/repository/order_repository.go -destination mock/mock_repository/order_repository_mock.go
mockgen -source domain/repository/order_states_repository.go -destination mock/mock_repository/order_states_repository_mock.go
mockgen -source domain/repository/order_type_repository.go -destination mock/mock_repository/order_type_repository_mock.go
mockgen -source domain/repository/project_cost_detail_repository.go -destination mock/mock_repository/project_cost_detail_repository_mock.go
mockgen -source domain/repository/quotation_detail_repository.go -destination mock/mock_repository/quotation_detail_repository_mock.go
mockgen -source domain/repository/quotation_history_repository.go -destination mock/mock_repository/quotation_history_repository_mock.go
mockgen -source domain/repository/quotation_items_repository.go -destination mock/mock_repository/quotation_items_repository_mock.go
mockgen -source domain/repository/quotations_repository.go -destination mock/mock_repository/quotations_repository_mock.go
mockgen -source domain/repository/acceptance_details_repository.go -destination mock/mock_repository/acceptance_details_repository.mock.go
mockgen -source domain/repository/order_history_repository.go -destination mock/mock_repository/order_history_repository_mock.go