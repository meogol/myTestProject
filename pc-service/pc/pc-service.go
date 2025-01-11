package pc

import (
	"errors"
	pcModel "meogol/pc-service/database/pc"
)

func createPc(r *Request) error {
	err := pcModel.Create(toModel(r))
	if err != nil {
		pcLogger.DPanicf("failed to save pc; Reason: %s", err)
		return err
	}
	return nil
}

func updatePc(r *Request) error {
	if r.Id.IsNone() {
		return errors.New("pc id is required")
	}

	err := pcModel.Update(r.Id.Get(), toModel(r))
	if err != nil {
		pcLogger.DPanicf("failed to update pc; Reason: %s", err)
		return err
	}
	return nil
}

func deletePc(id int) error {
	err := pcModel.Delete(id)
	if err != nil {
		pcLogger.DPanicf("failed to delete pc; Reason: %s", err)
		return err
	}
	return nil
}

func getPc(id int) (*pcModel.Model, error) {
	pc, err := pcModel.Get(id)
	if err != nil {
		pcLogger.DPanicf("failed to get pc; Reason: %s", err)
		return nil, err
	}
	return pc, nil
}

func toModel(pc *Request) *pcModel.Model {
	return &pcModel.Model{
		Name:        pc.Name,
		Description: pc.Description,
		Processor:   pc.Processor,
		VideoCard:   pc.VideoCard,
	}
}
