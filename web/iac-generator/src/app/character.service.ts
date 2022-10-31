import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { Character } from './character';
import { environment } from '../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class CharacterService {
  private characterUrl = `http://${environment.apiUrl}:${environment.apiPort}/apcg/api/v1/character`;

  constructor(private http: HttpClient) {}

  get(): Observable<Character> {
    return this.http.get<any>(this.characterUrl).pipe(
      map((response) => {
        console.log(response);
        let c: Character = { name: response.name };
        return c;
      })
    );
  }
}
