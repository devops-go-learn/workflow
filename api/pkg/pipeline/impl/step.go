package impl

import (
	"context"

	"github.com/infraboard/keyauth/client/session"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/grpc/gcontext"
	"github.com/infraboard/workflow/api/pkg/pipeline"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (i *impl) QueryStep(ctx context.Context, req *pipeline.QueryStepRequest) (
	*pipeline.StepSet, error) {
	listKey := pipeline.EtcdStepPrefix()
	i.log.Infof("list etcd step resource key: %s", listKey)
	resp, err := i.client.Get(ctx, listKey, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	ps := pipeline.NewStepSet()
	for index := range resp.Kvs {
		// 解析对象
		ins, err := pipeline.LoadStepFromBytes(resp.Kvs[index].Value)
		if err != nil {
			i.log.Error(err)
			continue
		}
		ps.Add(ins)
	}
	return ps, nil
}

func (i *impl) DescribeStep(ctx context.Context, req *pipeline.DescribeStepRequest) (
	*pipeline.Step, error) {
	in, err := gcontext.GetGrpcInCtx(ctx)
	if err != nil {
		return nil, err
	}

	if req.Namespace == "" {
		tk := session.S().GetToken(in.GetAccessToKen())
		if tk == nil {
			return nil, exception.NewUnauthorized("token required")
		}
		req.Namespace = tk.Namespace
	}

	descKey := pipeline.StepObjectKey(req.Key)
	i.log.Infof("describe etcd step resource key: %s", descKey)
	resp, err := i.client.Get(ctx, descKey)
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, exception.NewNotFound("step %s not found", req.Key)
	}

	if resp.Count > 1 {
		return nil, exception.NewInternalServerError("step find more than one: %d", resp.Count)
	}

	ins := pipeline.NewDefaultStep()
	for index := range resp.Kvs {
		// 解析对象
		ins, err = pipeline.LoadStepFromBytes(resp.Kvs[index].Value)
		if err != nil {
			i.log.Error(err)
			continue
		}
	}
	return ins, nil
}

func (i *impl) DeleteStep(ctx context.Context, req *pipeline.DeleteStepRequest) (
	*pipeline.Step, error) {
	descKey := pipeline.StepObjectKey(req.Id)
	i.log.Infof("delete etcd pipeline resource key: %s", descKey)
	resp, err := i.client.Delete(ctx, descKey, clientv3.WithPrevKV())
	if err != nil {
		return nil, err
	}

	if resp.Deleted == 0 {
		return nil, exception.NewNotFound("step %s not found", req.Id)
	}

	ins := pipeline.NewDefaultStep()
	for index := range resp.PrevKvs {
		// 解析对象
		ins, err = pipeline.LoadStepFromBytes(resp.PrevKvs[index].Value)
		if err != nil {
			i.log.Error(err)
			continue
		}
		ins.ResourceVersion = resp.Header.Revision
	}
	return ins, nil
}
