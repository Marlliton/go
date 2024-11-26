package examsvc

func (es *ExamService) Delete(id string) error {
	_, err := es.repo.Get(id)
	if err != nil {
		return err
	}

	return es.repo.Delete(id)
}
