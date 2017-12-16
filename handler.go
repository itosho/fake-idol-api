/**
 * @apiDefine IdolResource
 * @apiSuccess {String} id アイドルID
 * @apiSuccess {String} name 名前
 * @apiSuccess {Number} age 年齢
 * @apiSuccess {String} profile プロフィール情報
 * @apiSuccess {String} created_at 作成日時（ISO8601形式）
 * @apiSuccess {String} updated_at 更新日時（ISO8601形式）
 */

/**
 * @apiDefine IdolParam
 * @apiParam {String} [name] 名前
 * @apiParam {Number} [age=17] 年齢
 * @apiParam {String} [profile] プロフィール情報
 */
package main

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

/**
 * @api {get} /idols/ 1.アイドルリスト取得API
 * @apiName GetAllIdols
 * @apiGroup Idol
 * @apiDescription 全アイドル情報のリストを返す。
 *
 * @apiSuccess {Object[]} idols アイドルのリスト情報
 * @apiSuccess {String} idols.id アイドルID
 * @apiSuccess {String} idols.name 名前
 * @apiSuccess {Number} idols.age 年齢
 * @apiSuccess {String} idols.profile プロフィール情報
 * @apiSuccess {String} idols.created_at 作成日時（ISO8601形式）
 * @apiSuccess {String} idols.updated_at 更新日時（ISO8601形式）
 * @apiSuccessExample {json} Success-Response:
 *     HTTP/1.1 200 OK
 *     [
 *       {
 *         "id": 1,
 *         "name": "松村沙友理",
 *         "age": 25,
 *         "profile": "乃木坂46のメンバーの一人。",
 *         "created_at": "2017-12-01T10:10:10Z",
 *         "updated_at": "2017-12-01T10:10:10Z"
 *       },
 *       {
 *         "id": 2,
 *         "name": "小林由依",
 *         "age": 18,
 *         "profile": "欅坂46のメンバーの一人。",
 *         "created_at": "2017-12-02T10:10:10Z",
 *         "updated_at": "2017-12-02T10:10:10Z"
 *       }
 *     ]
 */
func (i *Impl) GetAllIdols(w rest.ResponseWriter, r *rest.Request) {
	idols := []Idol{}
	i.DB.Find(&idols)
	w.WriteJson(&idols)
}

/**
 * @api {get} /idols/:id 2.アイドル取得API
 * @apiName GetIdol
 * @apiGroup Idol
 * @apiDescription 指定したアイドル情報を返す。
 *
 * @apiUse IdolResource
 *
 * @apiError NotFound 指定したIDが存在しない場合
 * @apiErrorExample {json} Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error_message": "Resource not found"
 *     }
 */
func (i *Impl) GetIdol(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	idol := Idol{}
	if i.DB.First(&idol, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&idol)
}

/**
 * @api {post} /idols 3.アイドル新規作成API
 * @apiName CreateIdol
 * @apiGroup Idol
 * @apiDescription アイドルを新規に作成する。作成したアイドル情報を返す。
 *
 * @apiUse IdolParam
 * @apiParamExample {json} Request-Example:
 *     {
 *       "name": "西野七瀬",
 *       "age": 23,
 *       "profile": "乃木坂46随一の人気を誇るメンバー。"
 *     }
 *
 * @apiUse IdolResource
 *
 * @apiError BadRequest リクエストパラメーターが正しくない場合
 */
func (i *Impl) PostIdol(w rest.ResponseWriter, r *rest.Request) {
	idol := Idol{}
	if err := r.DecodeJsonPayload(&idol); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := i.DB.Save(&idol).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&idol)
}

/**
 * @api {put} /idols/:id 4.アイドル更新API
 * @apiName UpdateIdol
 * @apiGroup Idol
 * @apiDescription 指定したアイドル情報を更新する。更新したアイドル情報を返す。
 *
 * @apiUse IdolParam
 *
 * @apiUse IdolResource
 *
 * @apiError NotFound 指定したIDが存在しない場合
 * @apiError BadRequest リクエストパラメーターが正しくない場合
 */
func (i *Impl) PutIdol(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	idol := Idol{}
	if i.DB.First(&idol, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Idol{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if updated.Name != "" {
		idol.Name = updated.Name
	}
	if updated.Age > 0 {
		idol.Age = updated.Age
	}
	if updated.Profile != "" {
		idol.Profile = updated.Profile
	}

	if err := i.DB.Save(&idol).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&idol)
}

/**
 * @api {delete} /idols/:id 5.アイドル削除API
 * @apiName DeleteIdol
 * @apiGroup Idol
 * @apiDescription 指定したアイドル情報を削除する。成功した場合はレスポンスなし。
 *
 * @apiSuccessExample {json} Success-Response:
 *     HTTP/1.1 204 No Content
 *
 * @apiError NotFound 指定したIDが存在しない場合
 */
func (i *Impl) DeleteIdol(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	idol := Idol{}
	if i.DB.First(&idol, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&idol).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
