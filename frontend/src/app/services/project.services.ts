import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {Observable, map} from "rxjs";
import {IRequest} from "../models/request.model";
import {ConfigurationService} from "./configuration.services";
import { environment } from '../../environments/environment';
import { Page } from "ngx-pagination";

@Injectable({
    providedIn: 'root'
})
export class ProjectServices {
  urlPath = ""

  constructor(private http: HttpClient, private configurationService: ConfigurationService) {
    this.urlPath = configurationService.getValue("pathUrl")
  }


  getAll(page: number, searchName: String): Observable<IRequest>{
    return this.http.get<IRequest>('http://'+ this.urlPath +'/api/v1/connector/projects?' +
      'limit=10&page='+page + '&search=' + searchName)

    // TODO Написать запрос на получение всех проектов, учесть пагинацию, поиск
  }

  // @ts-ignore
  addProject(key: String): Observable<IRequest>{
    // TODO Написать запрос на добавление проета в БД. Добавление происходит по ключу проекта
  }

  // @ts-ignore
  deleteProject(id: Number): Observable<IRequest> {
    // TODO Написать запрос на удаление проекта. Удаление происходит по id проекта в БД.
  }
}
